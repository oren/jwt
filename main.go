package main

import "fmt"

func main() {
	token := createJWT()
	valid := validateJWT(token)
	fmt.Println("Created token", token)
	fmt.Println("Valid?", valid)
}

func createJWT() string {
	return "abcd"
}

func validateJWT(token string) bool {
	return true
}
