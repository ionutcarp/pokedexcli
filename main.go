package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	inputSlice := strings.Fields(text)
	for i := range inputSlice {
		inputSlice[i] = strings.ToLower(inputSlice[i])
	}
	return inputSlice
}

func main() {
	fmt.Println("Hello, World!")
}
