package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Khaz713/pokedexcli/internal/pokecache"
)

func GetPokemon(name string, cache *pokecache.Cache) (RespPokemon, error) {
	fullUrl := pokemonURL + name

	if val, ok := cache.Get(fullUrl); ok == true {
		var results RespPokemon
		err := json.Unmarshal(val, &results)
		if err != nil {
			return RespPokemon{}, err
		}

		return results, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return RespPokemon{}, err
	}
	if res.StatusCode == 200 {

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespPokemon{}, err
		}

		cache.Add(fullUrl, data)

		var results RespPokemon
		err = json.Unmarshal(data, &results)
		if err != nil {
			return RespPokemon{}, err
		}

		return results, nil
	}
	return RespPokemon{}, errors.New(res.Status)
}
