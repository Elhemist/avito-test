package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	serviceTable = "services"
	orderTable   = "orders"
	userTable    = "users"
)

type Config struct {
	Host     string
	Port     string
	Usename  string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Usename, cfg.DBName, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitTables(db *sqlx.DB) {
	query := `
	DROP TABLE IF EXISTS users;
    DROP TABLE IF EXISTS services;
    DROP TABLE IF EXISTS orders;
	CREATE TABLE users(
		id      serial              primary key,
		balance numeric             not null,
		reserve numeric DEFAULT 0   not null 
	);
	CREATE TABLE services(
		id      serial          primary key,
		name    varchar(255)    not null,
		price   numeric         not null
	);
	CREATE TABLE orders(
		id          serial      primary key,
		serviceId   INTEGER,
		userId      INTEGER,
		date        timestamp   default CURRENT_TIMESTAMP,
		status      boolean     default false
	);
	`
	_, _ = db.Exec(query)
}
func FillTables(db *sqlx.DB) {
	query := `
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	INSERT INTO users(balance) VALUES (1000);
	
	INSERT INTO services(name, price) VALUES ('First', 1);
	INSERT INTO services(name, price) VALUES ('Second', 10);
	INSERT INTO services(name, price) VALUES ('Third', 100);
	INSERT INTO services(name, price) VALUES ('Forth', 1000);
	INSERT INTO services(name, price) VALUES ('Fifth', 10000);
	INSERT INTO services(name, price) VALUES ('Sixth', 2);
	INSERT INTO services(name, price) VALUES ('Seventh', 3);
	INSERT INTO services(name, price) VALUES ('Eigth', 4);
	INSERT INTO services(name, price) VALUES ('Ninth', 5);
	INSERT INTO services(name, price) VALUES ('Tenth', 6);
	`
	_, _ = db.Exec(query)
}
