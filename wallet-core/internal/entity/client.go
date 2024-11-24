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
