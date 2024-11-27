package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading from file.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}