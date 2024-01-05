package money

import (
	"database/sql"
)

type Repository struct {
	Conn *sql.DB
}

func (repo *Repository) GetCustomerBalanceById(id string) Balance {
	query := "SELECT balance FROM accounts WHERE customer_id = $1"

	row := repo.Conn.QueryRow(query, id)

	var balance Balance
	row.Scan(&balance.Balance)

	return balance
}

type Balance struct {
	Balance float64 `json:"balance"`
}
