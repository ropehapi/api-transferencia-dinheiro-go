package database

import "github.com/ropehapi/api-transferencia-dinheiro-go/internal/entity"

type AccountInterface interface {
	GetAccountById(id string) (*entity.Account, error)
}
