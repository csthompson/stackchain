package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/csthompson/stackchain/internal/models"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var payload string

func init() {
	flag.StringVar(&payload, "payload", "", "Transaction Payload to sign")
}

type KeyPair struct {
	PrivKey string `json:"privKey"`
	PubKey  string `json:"pubKey"`
}

func main() {
	keyPairFile, err := os.Open("test_pair.json")
	if err != nil {
		fmt.Println(err)
	}
	defer keyPairFile.Close()
	keyPairB, _ := ioutil.ReadAll(keyPairFile)

	var keyPair KeyPair

	json.Unmarshal(keyPairB, &keyPair)

	pubKey, _ := hexutil.Decode(keyPair.PubKey)
	privKey, _ := hexutil.Decode(keyPair.PrivKey)

	//Sha256 checksum of the transaction payload
	checksum := sha256.Sum256([]byte(payload))
	//Sign the checksum
	signature, err := secp256k1.Sign(checksum[:], privKey)
	if err != nil {
		panic(err)
	}

	tx := models.Transaction{
		Signer:    pubKey,
		Checksum:  checksum[:],
		Signature: signature,
		Type: models.TxType{
			Resource: models.AccountTx,
			Action:   models.Create,
		},
		Body: []byte(payload),
	}

	txBytes, _ := json.Marshal(tx)

	encodedTx := base64.StdEncoding.EncodeToString(txBytes)

	fmt.Println(encodedTx)
}
