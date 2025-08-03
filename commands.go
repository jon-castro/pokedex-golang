package main

import (
	"errors"
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

	config.previousLocationUrl = &locations.Previous
	config.nextLocationUrl = &locations.Next

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}

	fmt.Print("Pokedex > ")

	return nil
}

func commandMapb(config *RequestConfig) error {
	if config.previousLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locations, err := GetLocationRequest(*config.previousLocationUrl)
	if err != nil {
		return err
	}

	config.nextLocationUrl = &locations.Next
	config.previousLocationUrl = &locations.Previous

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}

	fmt.Print("Pokedex > ")

	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displayes a help message\nexit: Exit the Podekex\nPokedex > ")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
