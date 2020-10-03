package txproc

import (
	"github.com/csthompson/stackchain/internal/models"
)

//Implements transaction routing based on the type of transaction
// We use simple switch statements instead of a radix tree or similar since routes are static

type Router struct {
}

//Route on the
func (self *Router) Route(tx models.Transaction) {
	txType := tx.Type
	resource := txType.Resource

	switch resource {
	case models.NamespaceTx:
		self.NamespaceRouter(tx)
	}
}

//Actions on namespace
func (self *Router) NamespaceRouter(tx models.Transaction) {
	action := tx.Type.Action
	//Create a new namespace
	if action == models.Create {

	}
}
