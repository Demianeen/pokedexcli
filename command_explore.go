package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location area provides")
	}
	if len(args) > 1 {
		return errors.New("you can provide only one location at a time")
	}

	locationArea := args[0]
	resp, err := cfg.pokeapiClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Found Pokemons:")
	for _, pokemon := range resp.Pokemon_encounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
