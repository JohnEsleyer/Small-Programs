package main

import "fmt"

func main() {
	inputStr := "racecar"

	if isPalindrome(inputStr) {
		fmt.Println(inputStr, "is a palindrome.")
	} else {
		fmt.Println(inputStr, "is not a palindrome.")
	}
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
