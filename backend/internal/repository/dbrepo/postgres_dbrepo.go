package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
)

const dbTimeout = 3 * time.Second

type PostgresDBRepo struct {
	DB *sql.DB
}

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.DB
}

func (db *PostgresDBRepo) SelectAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM public.users`
	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.Id,
			&user.PassportSeries,
			&user.PassportNumber,
			&user.Surname,
			&user.Name,
			&user.Patronymic,
			&user.Address,
		)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}

func (db *PostgresDBRepo) SelectUserByPassport(series, number int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT passport_series, passport_number, surname, name, patronymic, address 
FROM public.users
WHERE passport_series = $1 AND passport_number = $2`

	var user models.User
	err := db.DB.QueryRowContext(ctx, query, series, number).Scan(
		&user.PassportSeries,
		&user.PassportNumber,
		&user.Surname,
		&user.Name,
		&user.Patronymic,
		&user.Address,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDBRepo) InsertUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO public.users (passport_series, passport_number, surname, name, patronymic, address)
VALUES ($1, $2, $3, $4, $5, $6) returning id;`

	var userId int
	err := db.DB.QueryRowContext(ctx, query,
		user.PassportSeries,
		user.PassportNumber,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
	).Scan(&userId)

	if err != nil {
		return err
	}

	if userId == 0 {
		return errors.New("user has not been created")
	}

	return nil
}
