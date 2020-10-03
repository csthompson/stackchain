package accountstore

import (
	"context"

	"github.com/csthompson/stackchain/internal/models"
)

// Account persistence
type AccountStore interface {
	GetAccountBySigner(ctx context.Context, signer []byte) (*models.Account, error)
	//GetAccountByName(ctx context.Context, name []byte) (models.Account, error)
	CreateAccount(ctx context.Context, account *models.Account) error
	//UpdateAccount(ctx context.Context, accoun *models.Account) error
}
