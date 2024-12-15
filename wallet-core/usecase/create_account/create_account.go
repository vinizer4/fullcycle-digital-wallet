package create_account

import (
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(
	accountGateway gateway.AccountGateway,
	clientGateway gateway.ClientGateway,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (createAccountUseCase *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := createAccountUseCase.ClientGateway.Get(input.ClientID)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = createAccountUseCase.accountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
