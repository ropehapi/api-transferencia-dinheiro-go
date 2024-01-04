package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/money"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/customer/{id}", money.GetCustomerBalance)
	r.Post("/transfer", money.TransferBetweenCustomers)

	err = http.ListenAndServe(":"+os.Getenv("WEB_SERVER_PORT"), r)
	if err != nil {
		panic(err)
	}
}
