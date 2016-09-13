package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out keys/app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in keys/app.rsa -pubout > keys/app.rsa.pub
)

var signKey []byte
var rsaPub []byte

func init() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}

	var err2 error
	rsaPub, err2 = ioutil.ReadFile(pubKeyPath)
	if err2 != nil {
		log.Fatal("Error reading public key")
		os.Exit(1)
	}

}

func main() {
	token := createJWT()
	// fmt.Println("Created token", token)
	validateJWT(token)
}

func createJWT() []byte {
	claims := jws.Claims{}
	claims.Set("AccessToken", "level1")
	signMethod := jws.GetSigningMethod("HS512")
	token := jws.NewJWT(claims, signMethod)
	byteToken, err := token.Serialize(signKey)
	if err != nil {
		log.Fatal("Error serializing the key. ", err)
		os.Exit(1)
	}

	return byteToken
}

func validateJWT(token []byte) {
	w, err := jws.ParseJWT(token)
	if err != nil {
		log.Fatal("Error parsing the token. ", err)
		os.Exit(1)
	}

	if err := w.Validate(rsaPub, crypto.SigningMethodRS512); err != nil {
		log.Fatal("Error validating the token. ", err)
		os.Exit(1)
	}
}
