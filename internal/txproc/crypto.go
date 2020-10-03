package txproc

import (
	"github.com/csthompson/stackchain/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
)

//Verify the cryptographic signature on the transaction
func VerifySignature(tx models.Transaction) bool {

	ecdsaPubKey, err := crypto.SigToPub(tx.Checksum, tx.Signature)
	if err != nil {
		return false
	}
	if string(crypto.FromECDSAPub(ecdsaPubKey)) == string(tx.Signer) {
		return true
	} else {
		return false
	}

}
