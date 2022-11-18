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

func (r *UserPostgres) CheckUser(User atest.User) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT * FROM %s WHERE userId=$1 ", "users")
	row := r.db.QueryRow(query, User.Id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
