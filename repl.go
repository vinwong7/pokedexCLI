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

		command, ok := getCommands()[first_word]
		if ok {
			command.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
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
	}
}
