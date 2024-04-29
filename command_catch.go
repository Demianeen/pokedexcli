package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you need to provide a pokemon name")
	}
	if len(args) > 1 {
		return errors.New("you can catch only one pokemon at a time")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(500 * time.Millisecond)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	cfg.caughtPokemons[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
