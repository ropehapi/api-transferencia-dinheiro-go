package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/entity"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/infra/database"
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/infra/webserver/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open("root:Pedroka0123@tcp(localhost:3306)/db_api_transferencia_dinheiro_go?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Account{})
	if err != nil {
		panic("failed to migrate database")
	}
	accountDB := database.NewAccount(db)
	accountHandler := handlers.NewAccountHandler(accountDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/account/{id}", accountHandler.GetAccountById)

	err = http.ListenAndServe(":"+os.Getenv("WEB_SERVER_PORT"), r)
	if err != nil {
		panic(err)
	}
}
