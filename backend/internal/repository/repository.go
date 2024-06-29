package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	SelectAllUsers() ([]*models.User, error)
	SelectUserByPassport(series, number int) (*models.User, error)
	InsertUser(user models.User) error
}
