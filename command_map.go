package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
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

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("No previous page")
	}	

	res, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
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