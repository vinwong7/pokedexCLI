package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words_list := strings.Fields(lowered_text)
	return words_list
}

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input_text := scanner.Text()
		clean_input := cleanInput(input_text)
		first_word := clean_input[0]

		command, ok := getCommands()[first_word]
		if ok {
			command.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
