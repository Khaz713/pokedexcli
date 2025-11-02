package main

import (
	"fmt"
)

func commandPokedex(_ []string, config *config) error {
	for _, key := range config.pokeapiClient.Pokedex.List() {
		fmt.Println("-", key)
	}
	return nil
}
