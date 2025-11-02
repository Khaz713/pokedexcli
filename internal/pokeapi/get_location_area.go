package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Khaz713/pokedexcli/internal/pokecache"
)

func GetLocationAreas(pageUrl *string, cache *pokecache.Cache) (RespLocationArea, error) {
	fullUrl := fmt.Sprintf("%s%s", baseURL, "/location-area")
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	if val, ok := cache.Get(fullUrl); ok == true {
		var results RespLocationArea
		err := json.Unmarshal(val, &results)
		if err != nil {
			return RespLocationArea{}, err
		}

		return results, nil
	}
	res, err := http.Get(fullUrl)
	if err != nil {
		return RespLocationArea{}, err
	}
	if res.StatusCode == 200 {

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationArea{}, err
		}

		cache.Add(fullUrl, data)

		var results RespLocationArea
		err = json.Unmarshal(data, &results)
		if err != nil {
			return RespLocationArea{}, err
		}

		return results, nil
	}
	return RespLocationArea{}, errors.New(res.Status)

}
