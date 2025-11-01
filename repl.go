package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	areaOffset int
}

func replStart() {
	config := &Config{areaOffset: 0}
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
			err := cliCommands[command[0]].callback(config)
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
	callback    func(config *Config) error
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
			description: "Displays map of 20 next areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays map of 20 previous areas",
			callback:    commandMapB,
		},
	}
}
