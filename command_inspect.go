package main

import "fmt"

func commandInspect(param []string, config *config) error {
	pokemon := param[1]

	if val, ok := config.pokeapiClient.Pokedex.Get(pokemon); ok {
		fmt.Println(val)
	} else {
		fmt.Printf("%s not found in pokedex\n", pokemon)
	}

	return nil
}
