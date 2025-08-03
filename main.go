package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	initialLocationUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	requestConfig := RequestConfig{
		previousLocationUrl: nil,
		nextLocationUrl:     &initialLocationUrl,
	}

	commands := map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display the next 20 map locations",
			callback: func() error {
				return commandMap(&requestConfig)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Display the next 20 map locations",
			callback: func() error {
				return commandMapb(&requestConfig)
			},
		},
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
			case "map":
				err := commands["map"].callback()
				if err != nil {
					return
				}
			case "mapb":
				err := commands["mapb"].callback()
				if err != nil {
					return
				}
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
				fmt.Print("Unknown command\nPokedex > ")
			}
		} else {
			fmt.Print("Pokedex > ")
		}
	}
}
