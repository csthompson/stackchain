package namespacestore

import (
	"context"

	"github.com/csthompson/stackchain/internal/models"
)

// Namespace persistence
type NamespaceStore interface {
	GetNamespace(ctx context.Context, name []byte) (models.Namespace, error)
	CreateNamespace(ctx context.Context, namespace *models.Namespace) error
	UpdateNamespace(ctx context.Context, namespace *models.Namespace) error
}
