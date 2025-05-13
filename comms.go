package main

import (
	"fmt"
	"internal/pokeApi"
	"os"
)

type config struct {
	previous string
	next     string
}

var Configs config = config{
	previous: "null",
	next:     "https://pokeapi.co/api/v2/location",
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      config
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Not closed")
}

func commandHelp() error {
	fmt.Printf("Usage: \n\n")
	s := getCommands()
	for command := range s {
		fmt.Printf("%s: %s\n", s[command].name, s[command].description)
	}
	return nil
}

func commandMap() error {
	locations, err := pokeApi.GetLocations(Configs.next)
	if err != nil {
		return err
	}
	updtConfigNext()
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapb() error {
	if Configs.previous == "null" || Configs.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := pokeApi.GetLocations(Configs.previous)
	if err != nil {
		return err
	}
	updtConfigPrev()
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	var commands map[string]cliCommand = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      Configs,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      Configs,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 locations on pokemon world",
			callback:    commandMap,
			config:      Configs,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the past 20 locations on pokemon world",
			callback:    commandMapb,
			config:      Configs,
		},
	}
	return commands
}

func updtConfigNext() error {
	next, previous, err := pokeApi.GetNextPrevious(Configs.next)
	if err != nil {
		return err
	}
	Configs.next = next
	Configs.previous = previous
	return nil
}

func updtConfigPrev() error {
	next, previous, err := pokeApi.GetNextPrevious(Configs.previous)
	if err != nil {
		return err
	}
	Configs.next = next
	Configs.previous = previous
	return nil
}
