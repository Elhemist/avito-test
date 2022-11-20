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
	var balance int
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id=$1 ", userTable)
	err := r.db.Get(&balance, query, order.User)
	if err != nil {
		return 0, err
	}
	var service atest.Service
	query = fmt.Sprintf("SELECT id,name, price FROM %s WHERE id=$1 ", serviceTable)
	err = r.db.Get(&service, query, order.Service)
	if balance-service.Price >= 0 {
		query = fmt.Sprintf("UPDATE %s SET balance =balance-%d, reserve =reserve+%d WHERE id=%d;", userTable, service.Price, service.Price, order.User)
		_, err = r.db.Exec(query)
	} else {
		return 0, err
	}
	var id int
	query = fmt.Sprintf("INSERT INTO %s (id, serviceId, userId) values ($1, $2, $3) RETURNING id", orderTable)
	row := r.db.QueryRow(query, order.Id, order.Service, order.User)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *OrderPostgres) RevenueOrder(order atest.Order) (int, error) {

	var res int
	query := fmt.Sprintf("SELECT reserve FROM %s WHERE id=$1 ", userTable)
	err := r.db.Get(&res, query, order.User)
	if err != nil {
		return 0, err
	}

	var servicePrice int
	query = fmt.Sprintf("SELECT price FROM %s WHERE id=$1 ", serviceTable)
	err = r.db.Get(&servicePrice, query, order.Service)
	if err != nil {
		return 0, err
	}

	var status bool
	query = fmt.Sprintf("SELECT status FROM %s WHERE id=$1 ", orderTable)
	err = r.db.Get(&status, query, order.Id)
	if err != nil {
		return 0, err
	}
	if status {
		return 0, err
	}

	tx, err := r.db.Begin()
	fmt.Println("nya3")
	if res-servicePrice >= 0 {
		res = res - servicePrice
		query = fmt.Sprintf("UPDATE %s SET reserve =%d WHERE id=%d;", userTable, res, order.User)
		_, err = tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		query = fmt.Sprintf("UPDATE %s SET status =%t WHERE id=%d;", orderTable, true, order.Id)
		_, err = tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		err = tx.Commit()
	} else {
		return 0, err
	}
	return order.Id, nil
}
