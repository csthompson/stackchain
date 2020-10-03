package statestore

import (
	"context"
)

// Stackchain State persistence
type StateStore interface {
	IncTxCount(ctx context.Context) error
	GetTxCount(ctx context.Context) (uint64, error)
}
