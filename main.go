package main

import (
	"time"

	"github.com/vinwong7/pokedexCLI/internal/pokeapi"
)

func main() {
	//replStart()
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	config := &config{
		pokeapiClient:  pokeClient,
		caught_pokemon: make(map[string]pokeapi.Pokestat),
	}
	replStart(config)
}
