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
	var userRes atest.User
	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", userTable)
	var res int
	err := r.db.Get(&userRes, query, order.User)
	query = fmt.Sprintf("SELECT reserve FROM %s WHERE id=$1 ", userTable)
	err = r.db.Get(&res, query, order.User)
	var service atest.Service
	query = fmt.Sprintf("SELECT id,name, price FROM %s WHERE id=$1 ", serviceTable)
	err = r.db.Get(&service, query, order.Service)
	if userRes.Balance-service.Price > 0 {
		userRes.Balance = userRes.Balance - service.Price
		res = res + service.Price
		query = fmt.Sprintf("UPDATE %s SET balance =%d reserve =%d WHERE id=%d;", userTable, userRes.Balance, res, userRes.Id)
		err = r.db.Get(userRes, query)
	}

	var id int
	query = fmt.Sprintf("INSERT INTO %s (id, serviceId, userId) values ($1, $2, $3) RETURNING id", orderTable)
	row := r.db.QueryRow(query, order.Id, order.Service, order.User)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
