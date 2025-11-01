package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataLocationArea struct {
	Results []LocationArea `json:"results"`
}
type LocationArea struct {
	Name string `json:"name"`
}

func GetLocationAreas(offset int) ([]LocationArea, error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d", offset)
	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var results DataLocationArea
	err = json.Unmarshal(data, &results)
	if err != nil {
		return nil, err
	}

	return results.Results, nil
}
