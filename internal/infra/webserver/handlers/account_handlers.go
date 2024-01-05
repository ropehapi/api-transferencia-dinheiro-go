package handlers

import (
	"encoding/json"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/infra/database"
	"net/http"
)

type AccountHandler struct {
	AccountDB database.AccountInterface
}

func NewAccountHandler(db database.AccountInterface) *AccountHandler {
	return &AccountHandler{
		AccountDB: db,
	}
}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	account, err := h.AccountDB.GetAccountById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
