package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	newScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		newScanner.Scan()
		text := newScanner.Text()
		text = strings.ToLower(text)
		words := strings.Fields(text)
		fmt.Println("Your command was:", words[0])
	}
}
