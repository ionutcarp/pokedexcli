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
	pokedex             map[string]pokeapi.Pokemon
	seenPokemon         map[string]struct{}
}
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore <location_name>",
			description: "Displays details about a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemnon_name>",
			description: "Attempts to catch a pokemnon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the Pokedex",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemnon_name>",
			description: "Lists details about a caught Pokemon",
			callback:    commandInspect,
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
		err := command.callback(cfg, words[1:]...)
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
