package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if _, seen := cfg.seenPokemon[pokemonName]; !seen {
		return fmt.Errorf("You haven't encountered %s so far, keep exploring...%v\n", pokemonName)
	}

	if _, owned := cfg.pokedex[pokemonName]; owned {
		return fmt.Errorf("You have already caught %s, check your pokedex!\n", pokemonName)
	}

	fmt.Printf("Throwing a Pokeball at %s..\n", pokemonName)
	threshold := int(0.6 * float32(pokemon.BaseExperience))
	attempt := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("You need to roll more than %d, to catch %s...\n", threshold, pokemonName)
	// We need a roll greater than 60% experience to catch
	if attempt > threshold {
		fmt.Printf("Congratulations!!! You have caught %s with a roll of %d.\n", pokemonName, attempt)
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s added to our pokedex:\n", pokemon.Name)
		for poke := range cfg.pokedex {
			fmt.Printf(" - %s", poke)
		}
		fmt.Println()
	} else {
		fmt.Printf("Better luck next time! You rolled %d.\n", attempt)
	}
	return nil
}
