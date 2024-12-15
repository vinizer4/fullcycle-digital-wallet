package database

import (
	"database/sql"
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	db.Exec("create table accounts (id varchar(255), client_id varchar(255), balance int, created_at datetime, updated_at datetime)")

	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("drop table clients")
	s.db.Exec("drop table accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	// Given
	s.db.Exec("insert into clients (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt,
	)

	account := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)

	// When
	accountDB, err := s.AccountDB.FindByID(account.ID)

	// Then
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.CreatedAt.In(time.Local), accountDB.CreatedAt.In(time.Local))
	s.Equal(account.UpdatedAt.In(time.Local), accountDB.UpdatedAt.In(time.Local))
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
	s.Equal(account.Client.CreatedAt.In(time.Local), accountDB.Client.CreatedAt.In(time.Local))
	s.Equal(account.Client.UpdatedAt.In(time.Local), accountDB.Client.UpdatedAt.In(time.Local))
}
