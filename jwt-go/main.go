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
	// eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJBY2Nlc3NUb2tlbiI6ImxldmVsMSIsIkN1c3RvbVVzZXJJbmZvIjp7Ik5hbWUiOiJqb3NoIiwiUm9sZSI6ImFkbWluIn0sImV4cCI6MTQ2NjQ2Njg2MX0.5gJQd7ql7vtfwcHm6ZHUxN3zmukQNUoYDdpBQ3-B0WKIHwFo6Arg0D-tiDWVaW8C4oAWBtNMohlJsBaOTbv5vg
}

func createJWT() string {
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims["AccessToken"] = "level1"
	token.Claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	token.Claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{"josh", "admin"}

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Fatal("Error signing the key")
		os.Exit(1)
	}

	return tokenString
}
