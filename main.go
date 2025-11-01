package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		input.Scan()
		inputString := input.Text()
		command := cleanInput(inputString)
		fmt.Println("Your command was:", command[0])
	}
}
