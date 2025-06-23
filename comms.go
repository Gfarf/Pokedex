package main

import (
	"errors"
	"fmt"
	"os"

	"internal/pokeApi"
)

type config struct {
	pokeapiClient    pokeApi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Not closed")
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Usage: \n\n")
	s := getCommands()
	for command := range s {
		fmt.Printf("%s: %s\n", s[command].name, s[command].description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	fmt.Println()
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	fmt.Println()
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	exploreResp, err := cfg.pokeapiClient.ExploreLocations(args[0])
	if err != nil {
		return err
	}
	fmt.Println()
	for _, pok := range exploreResp.Pokemon_encounters {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	var commands map[string]cliCommand = map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Get the next 20 locations on pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the past 20 locations on pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Return pokemons in the area",
			callback:    commandExplore,
		},
	}
	return commands
}
