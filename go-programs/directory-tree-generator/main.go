package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Recursively walk the directory tree and prin the names of all the files and directories.
	walk(cwd)
}

func walk(dir string) {
	// Get a list of all the files and directories and in the current directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the list of files and directories and print their names
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			walk(filepath.Join(dir, file.Name()))
		} else {
			fmt.Println("  ", file.Name())
		}
	}
}
