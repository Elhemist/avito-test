package repository

import (
	atest "github.com/Elhemist/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order atest.Order) (int, error)
}
type User interface {
	CheckUser(user atest.User) (int, error)
}

type Repository struct {
	Order
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
