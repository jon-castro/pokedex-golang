package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var finalSlice []string
	stripedText := strings.TrimSpace(strings.ToLower(text))
	finalSlice = strings.Fields(stripedText)

	return finalSlice
}
