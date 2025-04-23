package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words_list := strings.Fields(lowered_text)
	return words_list
}
