// A program that takes an input from the user
// and writes it to a file in the same directory

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get input string from the user
	fmt.Print("Enter string to write to file: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	// Create a new file in the same directory as the Go code file
	file, err := os.Create("./output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write input string to file
	_, err = file.WriteString(inputString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully wrote input string to file.")
}
