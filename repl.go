package main

import "strings"

func cleanInput(text string) []string {
	var output []string
	for _, word := range strings.Split(text, " ") {
		word = strings.TrimSpace(word)
		if len(word) > 0 {
			output = append(output, strings.ToLower(word))
		}

	}
	return output
}
