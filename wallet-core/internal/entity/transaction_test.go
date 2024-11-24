package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransactoin(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@j")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateNewTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@j")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 10000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
