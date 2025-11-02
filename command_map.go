package main

import (
	"fmt"

	"github.com/Khaz713/pokedexcli/internal/pokeapi"
)

func commandMapF(config *config) error {
	resp, err := pokeapi.GetLocationAreas(config.nextLocationURL, config.pokeapiClient.Cache)
	if err != nil {
		return err
	}
	config.nextLocationURL = resp.Next
	config.prevLocationURL = resp.Previous
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(config *config) error {
	if config.prevLocationURL == nil {
		fmt.Println("You're on the first page")
	} else {
		resp, err := pokeapi.GetLocationAreas(config.prevLocationURL, config.pokeapiClient.Cache)
		if err != nil {
			return err
		}
		config.nextLocationURL = resp.Next
		config.prevLocationURL = resp.Previous
		for _, area := range resp.Results {
			fmt.Println(area.Name)
		}
	}
	return nil

}
