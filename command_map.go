package main

import (
	"fmt"

	"github.com/Khaz713/pokedexcli/pokeApi"
)

func commandMap(config *Config) error {
	areas, err := pokeApi.GetLocationAreas(config.areaOffset)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	config.areaOffset += 20
	return nil
}
