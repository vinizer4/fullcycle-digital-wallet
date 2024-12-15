package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
