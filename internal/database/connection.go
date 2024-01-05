package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db_api_transferencia_dinheiro_go.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(conn *sql.DB) {
	stmt, err := conn.Prepare("CREATE TABLE IF NOT EXISTS accounts (customer_id BIGINT(19) NULL DEFAULT NULL,balance FLOAT NULL DEFAULT NULL)")
	if err != nil {
		panic(err)
	}

	stmt, err = conn.Prepare("INSERT INTO accounts (customer_id, balance) VALUES(1, 5000)")
	if err != nil {
		panic(err)
	}

	stmt.Exec()
}
