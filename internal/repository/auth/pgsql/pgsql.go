package pgsql

import (
	"database/sql"
	"errors"
)

var (
	ErrNoAuthorization = errors.New("no authorization")
)

type Repository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
