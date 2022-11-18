package repository

import "github.com/jmoiron/sqlx"

type Order interface {
}
type User interface {
}

type Repository struct {
	Order
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
