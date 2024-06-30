package dbrepo

import (
	"backend/internal/models"
	"context"
	"log"
)

func (db *PostgresDBRepo) SelectAllTasks() ([]*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT 
					id, user_id, task_name, started_at, finished_at, 
					EXTRACT( EPOCH FROM (finished_at - started_at))::INTEGER as duration
				FROM public.tasks`
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
