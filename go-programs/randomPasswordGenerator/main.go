package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	length := 12    // Specify the length of the password
	complexity := 3 // Specify the complexity of the password

	var password string
	for i := 0; i < length; i++ {
		password += generateRandomChar(complexity)
	}

	fmt.Println("Random Password:", password)
}

func generateRandomChar(complexity int) string {
	var char string
	switch complexity {
	case 1:
		char = "abcdefghijklmnopqrstuvwxyz"
	case 2:
		char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case 3:
		char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	}

	if n, err := rand.Int(rand.Reader, big.NewInt(int64(len(char)))); err != nil {
		panic(err)
	} else {
		return string(char[n.Int64()])
	}
}
