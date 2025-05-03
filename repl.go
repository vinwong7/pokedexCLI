package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vinwong7/pokedexCLI/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words_list := strings.Fields(lowered_text)
	return words_list
}

func replStart(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input_text := scanner.Text()
		clean_input := cleanInput(input_text)
		first_word := clean_input[0]
		var second_word string = ""
		if len(clean_input) > 1 {
			second_word = clean_input[1]
		}

		command, ok := getCommands()[first_word]
		if ok {
			command.callback(cfg, second_word)
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient  pokeapi.Client
	nextURL        *string
	previousURL    *string
	caught_pokemon map[string]pokeapi.Pokestat
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show a list of Pokemons found in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check the stats of a captured Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the names of all captured Pokemons",
			callback:    commandPokedex,
		},
	}
}
