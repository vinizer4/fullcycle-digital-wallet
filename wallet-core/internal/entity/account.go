package entity

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}

	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (acount *Account) Credit(amount float64) {
	acount.Balance += amount
	acount.UpdatedAt = time.Now()
}

func (acount *Account) Debit(amount float64) {
	acount.Balance -= amount
	acount.UpdatedAt = time.Now()
}
