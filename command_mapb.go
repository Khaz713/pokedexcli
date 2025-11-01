package main

import (
	"fmt"

	"github.com/Khaz713/pokedexcli/pokeApi"
)

func commandMapB(config *Config) error {
	if config.areaOffset < 40 {
		fmt.Println("You're on the first page")
	} else {
		config.areaOffset -= 20
		areas, err := pokeApi.GetLocationAreas(config.areaOffset - 20)
		if err != nil {
			return err
		}
		for _, area := range areas {
			fmt.Println(area.Name)
		}
	}
	return nil

}
