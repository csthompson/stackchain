package accountstore

import (
	"context"

	"github.com/csthompson/stackchain/internal/models"
	"github.com/dgraph-io/badger"
)

type AccountStore struct {
	Db *badger.DB
}

func (self *AccountStore) GetAccountBySigner(ctx context.Context, signer []byte) (*models.Account, error) {
	return nil, nil
}

func (self *AccountStore) CreateAccount(ctx context.Context, account *models.Account) error {
	return nil
}
