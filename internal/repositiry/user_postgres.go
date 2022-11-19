package repository

import (
	"fmt"

	atest "github.com/Elhemist/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}

}

func (r *UserPostgres) UpdateUser(user atest.User) (atest.User, error) {
	var userRes atest.User

	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", userTable)
	err := r.db.Get(&userRes, query, user.Id)
	if err != nil {
		return userRes, err
	}

	money := user.Balance + userRes.Balance
	query = fmt.Sprintf("UPDATE %s SET balance =%d WHERE id=%d;", userTable, money, user.Id)
	_, err = r.db.Exec(query)
	if err != nil {
		return userRes, err
	}

	query = fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", userTable)
	r.db.Get(&userRes, query, user.Id)

	return userRes, err
}
func (r *UserPostgres) CheckUser(user atest.User) (atest.User, error) {
	var userRes atest.User

	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1 ", "users")
	err := r.db.Get(&userRes, query, user.Id)
	return userRes, err
}
