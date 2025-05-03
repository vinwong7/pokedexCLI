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

	cache_data, ok := c.cache.Get(url)
	if !ok {

		//Send the GET request with http
		res, err := http.Get(url)
		if err != nil {
			return mapLocation{}, fmt.Errorf("error making request")
		}

		//Read the data with io
		res_data, err := io.ReadAll(res.Body)
		if err != nil {
			return mapLocation{}, fmt.Errorf("error reading response")
		}

		c.cache.Add(url, res_data)

		//Create an empty mapLocation struct, then unmarshal data onto it
		location_list := mapLocation{}
		if err := json.Unmarshal(res_data, &location_list); err != nil {
			return mapLocation{}, fmt.Errorf("unmarshal error")
		}

		//return the mapLocation struct to be used for the map/ mapb command
		return location_list, nil

	}

	//Create an empty mapLocation struct, then unmarshal data onto it
	cache_location_list := mapLocation{}
	if err := json.Unmarshal(cache_data, &cache_location_list); err != nil {
		return mapLocation{}, fmt.Errorf("unmarshal error")
	}

	//return the mapLocation struct to be used for the map/ mapb command
	return cache_location_list, nil

}
