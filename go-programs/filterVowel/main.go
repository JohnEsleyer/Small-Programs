package main

import (
	"fmt"
	"strings"
)

// A simple go program that prints all the vowels from the input string

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	vowels := []string{"a", "e", "i", "o", "u"}
	var vowelStr string

	for _, char := range strings.ToLower(input) {
		if contains(vowels, string(char)) && !contains(vowelStr, string(char)) {
			vowelStr += string(char)
		}
	}

	fmt.Println("Vowels in the string:", vowelStr)
}

func contains(slice interface{}, item interface{}) bool {
	switch slice.(type) {
	case []string:
		s := slice.([]string)
		for _, v := range s {
			if v == item {
				return true
			}
		}
	case string:
		s := slice.(string)
		for _, v := range s {
			if v == item {
				return true
			}
		}
	}
	return false
}
