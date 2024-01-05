package database

import (
	"github.com/ropehapi/api-transferencia-dinheiro-go/internal/entity"
	"gorm.io/gorm"
)

type Account struct {
	DB *gorm.DB
}

func NewAccount(db *gorm.DB) *Account {
	return &Account{DB: db}
}

func (a *Account) GetAccountById(id string) (*entity.Account, error) {
	var account entity.Account
	err := a.DB.First(&account, "id = ?", id).Error
	return &account, err
}
