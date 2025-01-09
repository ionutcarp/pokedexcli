package main

import (
	"github.com/ionutcarp/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
