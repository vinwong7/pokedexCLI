package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vinwong7/pokedexCLI/internal/pokecache"
)

var cache *pokecache.Cache = pokecache.NewCache(10 * time.Minute)

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

	cache_data, ok := cache.Get(url)
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

		cache.Add(url, res_data)

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

type pokemonLocation struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) PokemonLocations(location string) (pokemonLocation, error) {
	//Using location area here
	url := baseURL + "/location-area/" + location

	cache_data, ok := cache.Get(url)
	if !ok {

		//Send the GET request with http
		res, err := http.Get(url)
		if err != nil {
			return pokemonLocation{}, fmt.Errorf("error making request")
		}

		//Read the data with io
		res_data, err := io.ReadAll(res.Body)
		if err != nil {
			return pokemonLocation{}, fmt.Errorf("error reading response")
		}

		cache.Add(url, res_data)

		//Create an empty pokemonLocation struct, then unmarshal data onto it
		pokemon_list := pokemonLocation{}
		if err := json.Unmarshal(res_data, &pokemon_list); err != nil {
			return pokemonLocation{}, fmt.Errorf("unmarshal error")
		}

		//return the pokemonLocation struct to be used for the explore command
		return pokemon_list, nil

	}

	//Create an empty pokemonLocation struct, then unmarshal data onto it
	cache_pokemonList := pokemonLocation{}
	if err := json.Unmarshal(cache_data, &cache_pokemonList); err != nil {
		return pokemonLocation{}, fmt.Errorf("unmarshal error")
	}

	//return the pokemonLocation struct to be used for the explore command
	return cache_pokemonList, nil

}
