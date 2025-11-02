package main

import (
	"fmt"
	"math/rand"

	"github.com/Khaz713/pokedexcli/internal/pokeapi"
	"github.com/Khaz713/pokedexcli/internal/pokedex"
)

func commandCatch(param []string, config *config) error {
	resp, err := pokeapi.GetPokemon(param[1], config.pokeapiClient.Cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	if rand.Intn(resp.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", resp.Name)
	} else {
		fmt.Printf("%s was cought!\n", resp.Name)
		if _, ok := config.pokeapiClient.Pokedex.Get(resp.Name); ok != true {
			var types []string
			for _, val := range resp.Types {
				types = append(types, val.Type.Name)
			}
			config.pokeapiClient.Pokedex.Add(resp.Name, pokedex.Pokemon{
				Name:   resp.Name,
				Height: resp.Height,
				Weight: resp.Weight,
				Stats: struct {
					Hp             int
					Attack         int
					Defense        int
					SpecialAttack  int
					SpecialDefense int
					Speed          int
				}{Hp: resp.Stats[0].BaseStat, Attack: resp.Stats[1].BaseStat, Defense: resp.Stats[2].BaseStat, SpecialAttack: resp.Stats[3].BaseStat, SpecialDefense: resp.Stats[4].BaseStat, Speed: resp.Stats[5].BaseStat},
				Types: types,
			})
		} else {
			fmt.Printf("%s is already in your Pokedex\n", resp.Name)
		}
	}
	return nil
}
