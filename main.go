package main

import (
	"time"

	"github.com/Khaz713/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	replStart(cfg)
}
