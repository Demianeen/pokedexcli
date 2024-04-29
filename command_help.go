package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
	return nil
}
