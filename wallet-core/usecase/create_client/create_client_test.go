package create_client

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

func TestCreateClientUseCase_Execute(t *testing.T) {
	// Given
	mockedClientGateway := &ClientGatewayMock{}
	mockedClientGateway.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateClientUseCase(mockedClientGateway)

	// When
	output, err := usecase.Execute(CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j",
	})

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)
	mockedClientGateway.AssertExpectations(t)
	mockedClientGateway.AssertNumberOfCalls(t, "Save", 1)
}
