package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

import "github.com/SermoDigital/jose/jws"

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out keys/app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in keys/app.rsa -pubout > keys/app.rsa.pub
)

var signKey []byte

func init() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}

}

func main() {
	token := createJWT()
	fmt.Println("Created token", token)
}

func createJWT() string {
	claims := jws.Claims{}
	claims.Set("AccessToken", "level1")
	signMethod := jws.GetSigningMethod("HS512")
	token := jws.NewJWT(claims, signMethod)
	byteToken, err := token.Serialize(signKey)
	if err != nil {
		log.Fatal("Error serializing the key. ", err)
		os.Exit(1)
	}

	return string(byteToken)
}
