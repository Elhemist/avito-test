package repository

import (
	"fmt"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}

}

func (r *OrderPostgres) CreateOrder(order atest.Order) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, serviceId, userId) values ($1, $2, $3) RETURNING id", orderTable)
	row := r.db.QueryRow(query, order.Id, order.Service, order.User)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
