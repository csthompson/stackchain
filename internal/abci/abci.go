package abci

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/csthompson/stackchain/internal/models"
	"github.com/csthompson/stackchain/internal/plugin/storage/badger/accountstore"
	"github.com/csthompson/stackchain/internal/plugin/storage/badger/statestore"
	"github.com/csthompson/stackchain/internal/txproc"
	"github.com/prometheus/common/log"

	"github.com/dgraph-io/badger"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type Application struct {
	DB *badger.DB
}

var _ abcitypes.Application = (*Application)(nil)

func NewApplication(db *badger.DB) Application {
	return Application{db}
}

func (self Application) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (self Application) SetOption(req abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{}
}

func (self Application) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	return abcitypes.ResponseDeliverTx{Code: 0}
}

func (self Application) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	//Establish context
	ctx := context.Background()

	txB, err := base64.StdEncoding.DecodeString(string(req.Tx))

	fmt.Println(string(txB))

	//Get the transaction
	var tx models.Transaction
	if err := json.Unmarshal(txB, &tx); err != nil {
		panic(err)
	}

	//Create new Badger-based account store
	accountStore := accountstore.AccountStore{
		Db: self.DB,
	}
	//Create new Badger-based state store
	stateStore := statestore.StateStore{
		Db: self.DB,
	}

	//Create tx validator
	validator := txproc.TxValidator{
		AccountStore: &accountStore,
		StateStore:   &stateStore,
	}

	//Check if we are on the genesis transaction

	//Validate the signer of the transaction and the permissions
	statusCode, err := validator.ValidateTx(ctx, tx)
	if err != nil {
		log.Error(err)
		return abcitypes.ResponseCheckTx{Code: 400}
	}

	//Anything other than status code 0 is invalid
	if statusCode == 0 {
		return abcitypes.ResponseCheckTx{Code: 0}
	} else {
		return abcitypes.ResponseCheckTx{Code: statusCode}
	}
}

func (self Application) Commit() abcitypes.ResponseCommit {
	return abcitypes.ResponseCommit{}
}

func (self Application) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{Code: 0}
}

func (self Application) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (self Application) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (self Application) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}
