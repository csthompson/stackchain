package models

type Permission int

//Permissions are defined slightly different to bitmask to the actions
// Permissions are defined as RCUD (CRUD out of order to signify access level)
// For example, modifying existing data is more dangerous than create a new record

const (
	None  Permission = iota
	R                //Read
	C                //Create
	RC               //Read,Create
	U                //Update
	RU               //Read,Update
	CU               //Create,Update
	RCU              //Read,Create,Update
	D                //Delete
	RD               //Read, Delete
	CD               //Create, Delete
	RCD              //Read, Create, Delete
	UD               //Update, Delete
	RUD              //Read, Update, Delete
	CUD              //Create, Udpate, Delete
	RCUD             //Read, Create, Update, Delete
	X                //Execute
	RX               //Read, Execute
	CX               //Create, Execute
	RCX              //Read, Create, Execute
	UX               //Update, Execute
	RUX              //Read, Update, Execute
	CUX              //Crate, Update, Execute
	RCUX             //Read, Create, Update, Execute
	DX               //Delete, Execute
	RDX              //Read, Delete, Execute
	CDX              //Create, Delete, Execute
	RCDX             //Read, Create, Delete, Execute
	UDX              //Update, Delete, Execute
	RUDX             //Read, Update, Delete, Execute
	CUDX             //Create, Update, Delete, Execute
	RCUDX            //Read, Create, Update, Delete, Execute
)

//An Account defines permissions and access within the chain
type Account struct {
	Name    string //A friendly name for the account
	Address []byte //The public key of the account

	//Permission levels
	Permissions map[Resource]Permission
}

//Use bitmasking to check the account has permission to do the transaction
func (self *Account) CheckPermission(tx Transaction) bool {
	if permLevel, ok := self.Permissions[tx.Type.Resource]; ok {
		if byte(permLevel)&byte(tx.Type.Action) == byte(tx.Type.Action) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
