package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Printf("- %v\n", location.Name)
	}
	fmt.Println()
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapBackwards(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you are on the first page, so you can't go backwards")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Previous location areas:")
	for _, location := range resp.Results {
		fmt.Printf("- %v\n", location.Name)
	}
	fmt.Println()
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
