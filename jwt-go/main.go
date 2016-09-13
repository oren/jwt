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
	SecretKey := "WOW,MuchShibe,ToDogge"
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		os.Exit(1)
	}
}

func main() {
	token, err := createJWT()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created token", token)

	// validateJWT(token)
}

func createJWT() (string, error) {
	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["userid"] = "1234"
	// Expire in 5 mins
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created token", tokenString)
	return tokenString, nil
}

// func createJWT() (string, error) {
// 	type MyCustomClaims struct {
// 		Foo string `json:"foo"`
// 		jwt.StandardClaims
// 	}

// 	// Create the Claims
// 	claims := MyCustomClaims{
// 		"bar",
// 		jwt.StandardClaims{
// 			ExpiresAt: 15000,
// 			Issuer:    "test",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString(signKey)

// 	return ss, err
// }

// func validateJWT(tokenString string) {
// 	type MyCustomClaims struct {
// 		Foo string `json:"foo"`
// 		jwt.StandardClaims
// 	}

// 	// sample token is expired.  override time so it parses as valid
// 	at(time.Unix(0, 0), func() {
// 		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("AllYourBase"), nil
// 		})

// 		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
// 			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	})
// }
