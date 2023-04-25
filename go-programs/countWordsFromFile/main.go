package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Initialize a counter for the number of words
	count := 0

	// Iterate over each word in the file and increent the counter
	fmt.Printf("Total number of words: %d\n", count)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
