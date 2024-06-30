package dbrepo

import (
	"backend/internal/models"
	"context"
	"errors"
	"log"
	"strconv"
)

func (db *PostgresDBRepo) SelectAllTasks() ([]*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT 
					id, user_id, task_name, started_at, finished_at, 
					EXTRACT( EPOCH FROM (finished_at - started_at))::INTEGER as duration
				FROM public.tasks
				WHERE finished_at IS NOT NULL;
`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.Id,
			&task.UserId,
			&task.TaskName,
			&task.StartedAt,
			&task.FinishedAt,
			&task.Duration,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (db *PostgresDBRepo) StartTask(task models.RequestNewTask) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO public.tasks (user_id, task_name, started_at) VALUES ($1, $2, now()) returning id;`

	var newTaskId int
	err := db.DB.QueryRowContext(ctx, query,
		task.UserId,
		task.TaskName,
	).Scan(&newTaskId)

	if err != nil {
		return 0, err
	}

	return newTaskId, nil
}

func (db *PostgresDBRepo) FinishTask(taskId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE public.tasks SET finished_at = now() WHERE id = $1;`

	result, err := db.DB.ExecContext(ctx, query, taskId)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("task with id " + strconv.Itoa(taskId) + " not found")
	}

	return nil
}

func (db *PostgresDBRepo) DeleteTask(taskId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM public.tasks WHERE id = $1;`

	result, err := db.DB.ExecContext(ctx, query, taskId)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("task with id " + strconv.Itoa(taskId) + " not found")
	}

	return nil
}
