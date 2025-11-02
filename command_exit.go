package main

import (
	"fmt"
	"os"
)

func commandExit(_ []string, config *config) error {
	_ = config
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
