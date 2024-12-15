package create_transaction

import (
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (mockedAccountGateway *AccountGatewayMock) Save(account *entity.Account) error {
	args := mockedAccountGateway.Called(account)
	return args.Error(0)
}

func (mockedAccountGateway *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := mockedAccountGateway.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	// Given
	client1, _ := entity.NewClient("client1", "j@j")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "j@j")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	usecase := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)

	// When
	output, err := usecase.Execute(&inputDto)

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.TransactionID)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
