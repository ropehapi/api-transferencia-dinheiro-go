package money

import (
	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/database"
	"net/http"
)

type TransferRequest struct {
	UserIdFrom string  `json:"user_id_from"`
	UserIdTo   string  `json:"user_id_to"`
	Amount     float64 `json:"amount"`
}

func GetCustomerBalance(w http.ResponseWriter, r *http.Request) {
	conn, err := database.GetConn()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()

	repo := Repository{Conn: conn}
	repo.GetCustomerBalanceById(chi.URLParam(r, "id"))
}

func TransferBetweenCustomers(w http.ResponseWriter, r *http.Request) {

}
