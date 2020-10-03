package txproc

import (
	"context"
	"log"

	"github.com/csthompson/stackchain/internal/models"
	iaccountstore "github.com/csthompson/stackchain/internal/storage/accountstore"
	istatestore "github.com/csthompson/stackchain/internal/storage/statestore"
	"github.com/csthompson/stackchain/internal/types"
)

//TX validator provides validation facilities for incoming transactions and queries
type TxValidator struct {
	AccountStore iaccountstore.AccountStore
	StateStore   istatestore.StateStore
}

//Validate will perform a simple check on the transaction to verify the signer of the transaction and permissions
func (self *TxValidator) ValidateTx(ctx context.Context, tx models.Transaction) (uint32, error) {
	//Validate the signature
	sigIsVerify := VerifySignature(tx)
	if !sigIsVerify {
		return types.ErrorBadSignature, nil
	}

	//If this is the first transaction, we do not validate the permissions
	// First transaction must be creation and persistence of a fully permissioned account
	txCount, err := self.StateStore.GetTxCount(ctx)
	if err != nil {
		return types.ErrorInternal, err
	}
	if txCount == 0 && tx.Type.Resource == models.AccountTx && tx.Type.Action == models.Create {
		log.Println("Genesis transaction")
		return types.StatusOk, nil
	} else {

		//Validate the account permissions
		account, err := self.AccountStore.GetAccountBySigner(ctx, tx.Signer)
		if err != nil {
			return types.ErrorUnauthorized, err
		}

		hasPermission := self.checkPermissions(account, tx)
		if !hasPermission {
			return types.ErrorUnauthorized, err
		}

		return types.StatusOk, nil
	}
}

//Check the permissions of the account against the transaction resource/action
func (self *TxValidator) checkPermissions(account *models.Account, tx models.Transaction) bool {

	return false
}
