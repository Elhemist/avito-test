package repository

import (
	atest "github.com/Elhemist/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(atest.Order) (int, error)
}
type User interface {
	CheckUser(atest.User) (atest.User, error)
	UpdateUser(atest.User) (atest.User, error)
}

type Repository struct {
	Order
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
		User:  NewUserPostgres(db),
	}
}
