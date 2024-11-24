package entity

import (
	"errors"
	"github.com/google/uuid"

	"time"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt string
	UpdatedAt string
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	if err := client.Validate(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now().String()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) AddAccount(account *Account) error {
	if account.Client.ID != client.ID {
		return errors.New("account does not belong to client")
	}
	client.Accounts = append(client.Accounts, account)
	return nil
}
