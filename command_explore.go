package main

import (
	"fmt"

	"github.com/Khaz713/pokedexcli/internal/pokeapi"
)

func commandExplore(param []string, config *config) error {
	resp, err := pokeapi.GetLocationPokemon(param[1], config.pokeapiClient.Cache)
	if err != nil {
		return err
	}
	for _, val := range resp.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}
