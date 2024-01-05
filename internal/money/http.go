package money

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/database"
	"net/http"
)

func GetCustomerBalance(w http.ResponseWriter, r *http.Request) {
	conn, err := database.GetConnSQLite()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()
	database.AutoMigrate(conn)

	repo := Repository{Conn: conn}
	balance := repo.GetCustomerBalanceById(chi.URLParam(r, "id"))

	j, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(j)
}

type TransferRequest struct {
	UserIdFrom string  `json:"user_id_from"`
	UserIdTo   string  `json:"user_id_to"`
	Amount     float64 `json:"amount"`
}

func TransferBetweenCustomers(w http.ResponseWriter, r *http.Request) {

}
