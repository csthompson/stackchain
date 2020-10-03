package main

import (
	"flag"
)

var cmd string

var priKey string
var pubKey string

func init() {
	flag.StringVar(&cmd, "cmd", "ping", "Transaction command to run")
	flag.StringVar(&priKey, "priv", "-", "Private key")
	flag.StringVar(&pubKey, "pub", "-", "Public key")
}

func main() {

}
