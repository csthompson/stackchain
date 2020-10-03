package statestore

import (
	"context"
	"encoding/json"

	"github.com/csthompson/stackchain/internal/models"
	"github.com/dgraph-io/badger"
)

const (
	STATEKEY = "stackchain.application.state"
)

type StateStore struct {
	Db *badger.DB
}

//Increment the number of transaction executed on the chain by 1
func (self *StateStore) IncTxCount(ctx context.Context) error {
	err := self.Db.Update(func(txn *badger.Txn) error {
		var state models.State
		item, err := txn.Get([]byte(STATEKEY))
		//If the state is not set, persist a new one
		if err == badger.ErrKeyNotFound {
			newState := models.State{TxCount: 1}
			newStateB, _ := json.Marshal(newState)
			txn.Set([]byte(STATEKEY), newStateB)
		} else if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			if err := json.Unmarshal(val, &state); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		//Increment the transaction count
		state.TxCount++
		stateB, _ := json.Marshal(state)
		txn.Set([]byte(STATEKEY), stateB)
		return nil
	})
	return err
}

//Retrieve the number of transactions executed on the chain
func (self *StateStore) GetTxCount(ctx context.Context) (uint64, error) {
	var txCount uint64
	err := self.Db.View(func(txn *badger.Txn) error {
		var state models.State
		item, err := txn.Get([]byte(STATEKEY))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			if err := json.Unmarshal(val, &state); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		txCount = state.TxCount
		return nil
	})
	//If the err is that it doesnt exist, simple return 0
	if err == badger.ErrKeyNotFound {
		return 0, nil
	}
	return txCount, err
}
