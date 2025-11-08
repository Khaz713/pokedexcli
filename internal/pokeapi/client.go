package pokeapi

import (
	"net/http"
	"strings"
	"time"

	"github.com/Khaz713/pokedexcli/internal/pokecache"
	"github.com/Khaz713/pokedexcli/internal/pokedex"
)

type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
	Pokedex    *pokedex.Pokedex
}

func (c *Client) Do(line []rune, pos int) ([][]rune, int) {
	words := strings.Fields(strings.ToLower(string(line[:pos])))
	commands := []string{"help", "exit", "map", "mapb", "explore", "catch", "inspect", "pokedex"}
	suggestions := []string{}
	length := 0

	if len(words) == 0 {
		suggestions = commands
	} else if len(words) == 1 && line[pos-1] != ' ' {
		for _, command := range commands {
			if strings.HasPrefix(command, words[0]) {
				suggestions = append(suggestions, command)
			}
		}
		length = len([]rune(words[0]))
	} else if len(words) == 1 && line[pos-1] == ' ' {
		if words[0] == "inspect" {
			for _, key := range c.Pokedex.List() {
				suggestions = append(suggestions, key)
			}
		}
		// ADD AUTO COMPLETE FOR AREAS AND CATCHING POKEMON IN THE FUTURE
	} else {
		if words[0] == "inspect" {
			if line[len(line)-1] != ' ' {
				for _, key := range c.Pokedex.List() {
					if strings.HasPrefix(key, string(words[1])) {
						suggestions = append(suggestions, key)
					}
				}
				length = len([]rune(words[1]))
			}
		}
	}

	output := [][]rune{}
	for _, suggestion := range suggestions {
		output = append(output, []rune(suggestion[length:]))
	}
	return output, length
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache:   pokecache.NewCache(cacheInterval),
		Pokedex: pokedex.NewPokedex(),
	}
}
