package database

import (
	"database/sql"
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
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
