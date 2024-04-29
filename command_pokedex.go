package main

import "fmt"

func callbackPokedex(cfg *config, _ ...string) error {
	fmt.Println("Pokemons in Pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
