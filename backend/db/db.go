package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres password=1234 dbname=todo_app sslmode=disable")
	return err
}
