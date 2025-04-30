package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mapLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) MapLocations(pageURL *string) (mapLocation, error) {
	//Using location area here
	url := baseURL + "/location-area"

	//If there is a page available, use that page instead
	if pageURL != nil {
		url = *pageURL
	}

	//Send the GET request with http
	res, err := http.Get(url)
	if err != nil {
		return mapLocation{}, fmt.Errorf("error making request")
	}

	//Read the data with io
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return mapLocation{}, fmt.Errorf("error reading response")
	}

	//Create an empty mapLocation struct, then unmarshal data onto it
	location_list := mapLocation{}
	if err = json.Unmarshal(data, &location_list); err != nil {
		return mapLocation{}, fmt.Errorf("unmarshal error")
	}

	//return the mapLocation struct to be used for the map/ mapb command
	return location_list, nil
}
