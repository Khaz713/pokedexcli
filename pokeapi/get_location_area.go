package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(pageUrl *string) (RespLocationArea, error) {
	fullUrl := fmt.Sprintf("%s%s", baseURL, "/location-area")
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return RespLocationArea{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}
	var results RespLocationArea
	err = json.Unmarshal(data, &results)
	if err != nil {
		return RespLocationArea{}, err
	}

	return results, nil
}
