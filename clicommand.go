package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/vinwong7/pokedexCLI/internal/pokeapi"
)

func commandExit(c *config, s string) error {
	//Using os.exit to close out of the CLI
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	//Loop through the commands list and print out name and description
	//Using a function to generate the commands list rather than declare it as a variable
	//This avoids circular calling between this function and the list
	for _, v := range getCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}

	return nil
}

func commandMap(c *config, s string) error {

	//Runs the MapLocations method to generate the map location struct
	location, err := c.pokeapiClient.MapLocations(c.nextURL)
	if err != nil {
		return err
	}

	//Updates the map location struct with the newest URLs so rerunning the function runs the next page
	(*c).previousURL = location.Previous
	(*c).nextURL = location.Next

	//Loops through the map location struct results, which are the map locations
	for _, v := range location.Results {
		fmt.Printf("%v\n", v.Name)
	}
	return nil
}

func commandMapb(c *config, s string) error {

	//If the previous URL is nil, then we have not run past page 2, so just print statement
	if c.previousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	//Same as above, but using the previous URL
	location, err := c.pokeapiClient.MapLocations(c.previousURL)
	if err != nil {
		return err
	}

	(*c).previousURL = location.Previous
	(*c).nextURL = location.Next

	for _, v := range location.Results {
		fmt.Printf("%v\n", v.Name)
	}
	return nil
}

func commandExplore(c *config, s string) error {

	//Call the Pokemon Locations function, using the string input as the location to explore
	pokemons, err := c.pokeapiClient.PokemonLocations(s)
	if err != nil {
		return err
	}

	//Loop through the Pokemons from the returned data
	for _, v := range pokemons.PokemonEncounters {
		fmt.Printf("- %v\n", v.Pokemon.Name)
	}
	return nil
}

var caught_pokemon map[string]pokeapi.Pokestat = make(map[string]pokeapi.Pokestat)

func commandCatch(c *config, s string) error {

	stat, err := c.pokeapiClient.PokemonStats(s)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", s)
	base_chance := 1000 - stat.BaseExperience
	random_number := rand.Intn(1000)
	if base_chance >= random_number {
		fmt.Printf("%v was caught!\n", s)
		caught_pokemon[s] = stat

	} else {
		fmt.Printf("%v escaped!\n", s)
	}

	return nil
}

func commandInspect(c *config, s string) error {

	stat, ok := caught_pokemon[s]
	if !ok {
		println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %v\n", s)
		fmt.Printf("Height: %v\n", stat.Height)
		fmt.Printf("Weight: %v\n", stat.Weight)
		fmt.Println("Stats:")
		for _, v := range stat.Stats {
			fmt.Printf("	-%v: %v\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range stat.Types {
			fmt.Printf("	- %v\n", v.Type.Name)
		}
	}

	return nil
}

func commandPokedex(c *config, s string) error {

	fmt.Println("Your PokeDex:")
	for k := range caught_pokemon {
		fmt.Printf("= %v\n", k)
	}

	return nil
}
