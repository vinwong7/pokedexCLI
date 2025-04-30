package main

import (
	//"encoding/json"
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, v := range getCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}

	return nil
}

func commandMap(c *config) error {

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

func commandMapb(c *config) error {

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
