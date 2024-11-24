package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, 0.0, account.Balance)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithInvalidClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, 50.0, account.Balance)
}
