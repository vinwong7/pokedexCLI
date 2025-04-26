package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var location_config = config{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: nil,
}

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
	res, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("error making request")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response")
	}

	location := mapLocation{}
	if err = json.Unmarshal(data, &location); err != nil {
		return fmt.Errorf("unmarshal error")
	}

	(*c).Previous = location.Previous
	(*c).Next = location.Next

	for _, v := range location.Results {
		fmt.Printf("%v\n", v.Name)
	}
	return nil
}

func commandMapb(c *config) error {

	if c.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	res, err := http.Get(*c.Previous)
	if err != nil {
		return fmt.Errorf("error making request")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response")
	}

	location := mapLocation{}
	if err = json.Unmarshal(data, &location); err != nil {
		return fmt.Errorf("unmarshal error")
	}

	(*c).Previous = location.Previous
	(*c).Next = location.Next

	for _, v := range location.Results {
		fmt.Printf("%v\n", v.Name)
	}
	return nil
}
