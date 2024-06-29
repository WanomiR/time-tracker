package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.DB
}

func (db *PostgresDBRepo) AllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users`
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
			log.Println(err)
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}
