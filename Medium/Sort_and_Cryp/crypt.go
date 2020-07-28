package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	contra := `Cotorrisa`
	hash, err := bcrypt.GenerateFromPassword([]byte(contra), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
	fmt.Println(string(hash))
	bcrypt.CompareHashAndPassword(hash, []byte(contra))
}
