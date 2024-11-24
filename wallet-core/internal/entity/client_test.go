package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("Jane Doe Update", "j@j1.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe Update", client.Name)
	assert.Equal(t, "j@j1.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "")
	assert.NotNil(t, err)
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
