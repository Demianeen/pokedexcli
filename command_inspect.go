package main

import (
	"errors"
	"fmt"

	"github.com/Demianeen/pokedexcli/internal/pokeapi"
)

func inspectPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you need to provide a pokemon name")
	}
	if len(args) > 1 {
		return errors.New("you can view only one pokemon at a time")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return errors.New("you haven't called this pokemon yet")
	}
	inspectPokemon(pokemon)
	return nil
}
