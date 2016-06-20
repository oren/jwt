package main

import "fmt"
import "github.com/SermoDigital/jose/jws"

func main() {
	token := createJWT()
	valid := validateJWT(token)
	fmt.Println("Created token", token)
	fmt.Println("Valid?", valid)
}

func createJWT() string {
	var claims jws.Claims
	// claims.Set("AccessToken", "level1")

	signMethod := jws.GetSigningMethod("RS256")
	token := jws.NewJWT(claims, signMethod)
	fmt.Println("token", token)
	return "abcd"
}

func validateJWT(token string) bool {
	return true
}
