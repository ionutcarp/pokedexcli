package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	locationName := args[0]
	response, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("%s explored!\n", locationName)
	for _, det := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", det.Pokemon.Name)
		cfg.seenPokemon[det.Pokemon.Name] = struct{}{}
	}

	return nil
}
