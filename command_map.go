package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.GetLocationAreas()
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}