package main

import (
	"bufio"
	"fmt"
	"github.com/ionutcarp/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next map list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous map list",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", words[0])
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
