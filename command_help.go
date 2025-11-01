package main

import "fmt"

func commandHelp(config *Config) error {
	_ = config
	cliCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key := range cliCommands {
		fmt.Printf("%s: %s\n", cliCommands[key].name, cliCommands[key].description)
	}
	return nil
}
