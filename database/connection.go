package db

import (
	"context"
	"database/sql"
	config "web/app/api/config"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	SqlDb *sql.DB
}

var dbContext = context.Background()

func OpenConnection() (*sql.DB, error) {
	cfg := config.GetDbConnection()

	conn, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
