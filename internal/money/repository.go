package money

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	Conn *sql.DB
}

func (repo *Repository) GetCustomerBalanceById(id string) {
	query := "SELECT id, balance FROM accounts WHERE customer_id = $1"

	row := repo.Conn.QueryRow(query, id)

	fmt.Println(row.Scan())
}
