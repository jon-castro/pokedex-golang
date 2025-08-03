package main

import "strings"

func cleanInput(text string) []string {
	var finalSlice []string
	stripedText := strings.TrimSpace(strings.ToLower(text))
	finalSlice = strings.Fields(stripedText)

	return finalSlice
}
