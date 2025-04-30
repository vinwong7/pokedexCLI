package main

import (
	"github.com/vinwong7/pokedexCLI/internal/pokeapi"
)

func main() {
	//replStart()
	pokeClient := pokeapi.NewClient()
	config := &config{
		pokeapiClient: pokeClient,
	}
	replStart(config)
}
