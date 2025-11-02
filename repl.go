package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Khaz713/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func replStart(cfg *config) {
	cliCommands := getCommands()
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		inputString := input.Text()
		command := cleanInput(inputString)
		if len(command) == 0 {
			continue
		}

		if _, ok := cliCommands[command[0]]; ok == true {

			err := cliCommands[command[0]].callback(command, cfg)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Printf("Unknown command: %s\n", command[0])
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(param []string, config *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Get all the pokemon in given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch the pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get all pokemons in your pokedex",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect the pokemon in your pokedex",
			callback:    commandInspect,
		},
	}
}
