package models

type Resource int

const (
	TypeTx      Resource = iota //A type resource identifies the strucutre of a specific data type within
	ObjectTx                    //An object conforms to a specific data type and is an individual record
	CodeTx                      //A section of deterministic code to be executed and can utilize objects in the same namespace
	NamespaceTx                 //Globally unique identifier that associates models, objects, and code
	AccountTx                   //Namespace originator can configure accounts within the namespace and assign permissions. Each account is tied to a public address derived from BIP44 HD seed
)

type Action int

const (
	Read Action = 1 << iota
	Create
	Update
	Delete
	Execute
)

type TxType struct {
	Resource Resource
	Action   Action
}

type Transaction struct {
	Signer    []byte // The address of the originator
	Signature []byte // The signature of the transaction payload
	Checksum  []byte //Checksum of the transaction payload
	Type      TxType //The type of transaction
	Body      []byte //The body of the transaction
}
