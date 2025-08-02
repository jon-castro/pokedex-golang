package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display the application help screen",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		cleanedText := cleanInput(scanner.Text())
		if len(cleanedText) > 0 {
			switch cleanedText[0] {
			case "help":
				err := commands["help"].callback()
				if err != nil {
					return
				}
			case "exit":
				err := commands["exit"].callback()
				if err != nil {
					return
				}
			default:
				fmt.Println("Unknown command\nPokedex > ")
			}
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

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displayes a help message\nexit: Exit the Podekex")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
