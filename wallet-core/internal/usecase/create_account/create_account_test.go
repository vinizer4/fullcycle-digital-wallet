package create_account

import (
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (mockedClientGateway *ClientGatewayMock) Save(client *entity.Client) error {
	args := mockedClientGateway.Called(client)
	return args.Error(0)
}

func (mockedClientGateway *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := mockedClientGateway.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	// Given
	client, _ := entity.NewClient("John Doe", "j@j")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", mock.Anything).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateAccountUseCase(accountMock, clientMock)

	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	// When
	output, err := usecase.Execute(inputDto)

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
