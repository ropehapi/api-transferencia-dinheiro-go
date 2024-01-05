package entity

import (
	"errors"
	"github.com/ropehapi/api-transferencia-dinheiro-go/pkg/entity"
	"time"
)

var (
	ErrIdIsRequired   = errors.New("id is required")
	ErrBalanceInvalid = errors.New("balance is invalid")
)

type Account struct {
	ID        entity.ID `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(balance float64) (*Account, error) {
	account := &Account{
		ID:        entity.NewID(),
		Balance:   balance,
		CreatedAt: time.Now(),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Validate() error {
	if a.ID.String() == "" {
		return ErrIdIsRequired
	}

	if a.Balance < 0 {
		return ErrBalanceInvalid
	}

	return nil
}
