package main

import "github.com/adnant1/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}

	startRepl(&cfg)
}
