package common

import (
	"io/ioutil"
	"log"
)

const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"

	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath string = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey string
)

func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

}
