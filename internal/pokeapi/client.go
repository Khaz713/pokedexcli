package pokeapi

import (
	"net/http"
	"time"

	"github.com/Khaz713/pokedexcli/internal/pokecache"
	"github.com/Khaz713/pokedexcli/internal/pokedex"
)

type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
	Pokedex    *pokedex.Pokedex
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
