package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Khaz713/pokedexcli/internal/pokecache"
)

func GetLocationPokemon(area string, cache *pokecache.Cache) (RespLocationPokemon, error) {
	fullUrl := locationURL + area

	if val, ok := cache.Get(fullUrl); ok == true {
		var results RespLocationPokemon
		err := json.Unmarshal(val, &results)
		if err != nil {
			return RespLocationPokemon{}, err
		}

		return results, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return RespLocationPokemon{}, err
	}
	if res.StatusCode == 200 {

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationPokemon{}, err
		}

		cache.Add(fullUrl, data)

		var results RespLocationPokemon
		err = json.Unmarshal(data, &results)
		if err != nil {
			return RespLocationPokemon{}, err
		}

		return results, nil
	}
	return RespLocationPokemon{}, errors.New(res.Status)
}
