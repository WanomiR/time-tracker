package dbrepo

import (
	"database/sql"
	"time"
)

const dbTimeout = 3 * time.Second

type PostgresDBRepo struct {
	DB *sql.DB
}

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.DB
}
