package main

import (
	"fmt"

	"github.com/adnant1/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	res, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}