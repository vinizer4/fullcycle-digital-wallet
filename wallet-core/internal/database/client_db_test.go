package database

import (
	"database/sql"
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("create table clients (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("drop table clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSaveClient() {
	client, _ := entity.NewClient("John", "j@j.com")
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGetClient() {
	client, _ := entity.NewClient("John", "j@j.com")
	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
