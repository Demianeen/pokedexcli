package main

import (
	"time"

	"github.com/Demianeen/pokedexcli/internal/pokeapi"
)

type config struct {
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	pokeapiClient       pokeapi.Client
	caughtPokemons      map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:  pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
