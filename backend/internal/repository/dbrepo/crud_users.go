package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
)

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
			&user.Passport,
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

func (db *PostgresDBRepo) SelectUserByPassport(passport string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM public.users WHERE passport = $1`

	var user models.User
	err := db.DB.QueryRowContext(ctx, query, passport).Scan(
		&user.Id,
		&user.Passport,
		&user.Surname,
		&user.Name,
		&user.Patronymic,
		&user.Address,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user with passport " + passport + " not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDBRepo) SelectUserById(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM public.users WHERE id = $1`

	var user models.User
	err := db.DB.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Passport,
		&user.Surname,
		&user.Name,
		&user.Patronymic,
		&user.Address,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user with id " + strconv.Itoa(id) + " not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDBRepo) InsertUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO public.users (passport, surname, name, patronymic, address)
VALUES ($1, $2, $3, $4, $5);`

	_, err := db.DB.ExecContext(ctx, query,
		user.Passport,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDBRepo) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE public.users SET surname = $1, name = $2, patronymic = $3, address = $4 WHERE passport = $5;`

	result, err := db.DB.ExecContext(ctx, query,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
		user.Passport,
	)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("user with passport " + user.Passport + " not found")
	}

	return nil
}

func (db *PostgresDBRepo) DeleteUserByPassport(passport string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM public.users WHERE passport = $1;`

	result, err := db.DB.ExecContext(ctx, query, passport)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("user with passport " + passport + " not found")
	}

	return nil
}
