package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandMap(config *RequestConfig) error {
	locations, err := GetLocationRequest(*config.nextLocationUrl)
	if err != nil {
		return err
	}

	//TODO: DEBUG
	fmt.Println("Locations:", locations)

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}
	config.previousLocationUrl = config.nextLocationUrl
	config.nextLocationUrl = &locations.Next

	fmt.Print("Pokedex > ")

	return nil
}

func commandMapb(config *RequestConfig) error {
	if config.previousLocationUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := GetLocationRequest(*config.previousLocationUrl)
	if err != nil {
		return err
	}

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}
	config.nextLocationUrl = config.previousLocationUrl
	config.previousLocationUrl = &locations.Previous

	fmt.Print("Pokedex > ")

	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displayes a help message\nexit: Exit the Podekex\n")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
