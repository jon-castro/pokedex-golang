package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		cleanedText := cleanInput(scanner.Text())
		if len(cleanedText) > 0 {
			fmt.Printf("Your command was: %s\nPokedex > ", cleanedText[0])
		} else {
			fmt.Print("Pokedex > ")
		}
	}
}

func cleanInput(text string) []string {
	var finalSlice []string
	stripedText := strings.TrimSpace(strings.ToLower(text))
	finalSlice = strings.Fields(stripedText)

	return finalSlice
}
