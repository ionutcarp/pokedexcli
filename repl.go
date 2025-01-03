package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		fmt.Println("Your command was:", words[0])
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
