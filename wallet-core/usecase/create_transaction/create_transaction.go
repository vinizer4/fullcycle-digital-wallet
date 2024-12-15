package create_transaction

import (
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/gateway"
	"github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (createTransactionUseCase *CreateTransactionUseCase) Execute(inputDTO *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := createTransactionUseCase.AccountGateway.FindByID(inputDTO.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := createTransactionUseCase.AccountGateway.FindByID(inputDTO.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, inputDTO.Amount)
	if err != nil {
		return nil, err
	}

	err = createTransactionUseCase.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		TransactionID: transaction.ID,
	}, nil
}
