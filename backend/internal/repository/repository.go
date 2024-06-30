package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	SelectAllUsers() ([]*models.User, error)
	SelectUserByPassport(string) (*models.User, error)
	SelectUserById(int) (*models.User, error)
	InsertUser(models.User) error
	UpdateUser(models.User) error
	DeleteUserByPassport(string) error
	SelectAllTasks() ([]*models.Task, error)
	StartTask(models.RequestNewTask) (int, error)
	FinishTask(int) error
	DeleteTask(int) error
}
