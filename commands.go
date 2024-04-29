package main

type cliCommand struct {
	callback    func(*config, ...string) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			description: "Exits the program",
			callback:    callbackExit,
		},
		"help": {
			description: "Display a help message",
			callback:    callbackHelp,
		},
		"map": {
			description: "Lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			description: "List previous location areas",
			callback:    callbackMapBackwards,
		},
		"explore": {
			description: "List the pokemon in the location area",
			callback:    callbackExplore,
		},
		"catch": {
			description: "Attempts to catch the pokemon and adds it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			description: "Inspects one of caught pokemons",
			callback:    callbackInspect,
		},
		"pokedex": {
			description: "View all caught pokemons",
			callback:    callbackPokedex,
		},
	}
}
