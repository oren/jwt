package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out keys/app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in keys/app.rsa -pubout > keys/app.rsa.pub
)

func main() {
	token := createJWT()
	fmt.Println("Created token", token)
}

func createJWT() string {
	t := jwt.New(jwt.SigningMethodHS512)
	t.Claims["AccessToken"] = "level1"
	t.Claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	t.Claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{"josh", "admin"}

	signKey, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}

	tokenString, err := t.SignedString(signKey)

	return tokenString
}