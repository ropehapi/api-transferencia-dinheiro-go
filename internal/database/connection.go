package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetConn() (*sql.DB, error) {
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("CONN_STRING"))

	if err != nil {
		return nil, err
	}

	return conn, err
}
