package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/RussNavas/pokedex/internal/pokeapi"
)


func cleanInput(text string) []string{
		/*
		split the user's input into "words" based on whitespace. It should also
		lowercase the input and trim any leading or trailing whitespace.
		For example:
			hello world -> ["hello", "world"]
			Charmander Bulbasaur PIKACHU -> ["charmander", "bulbasaur", "pikachu"]
	*/

	pokemon := strings.Fields(text)
	
	for i, p := range pokemon{
		pokemon[i] = strings.ToLower(p)
	}


	return pokemon
}

var commands map[string]cliCommand
type cliCommand struct {
	name		string
	description string
	callback	func(*Config, []string) error
}


type Config struct{
	Next	 *string
	Previous *string
	Client	 *pokeapi.Client
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		"map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:	commandMapb,
		},
		"explore": {
		name:		"explore",
		description: "Displays the names of pokemon that can be encoutered in the area",
		callback:	commandExplore,
		},
	}
}


func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands = getCommands()
	for name, command := range commands{
		fmt.Printf("%v: %v\n", name, command.description)
	}
	return nil
}


func commandExit(config *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandMap(config *Config, args []string) error {
	locationsResp, err := config.Client.ListLocationAreas(config.Next)
	

	if err != nil {
		return err
	}

	config.Next = locationsResp.Next
	config.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results{
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(config *Config, args []string) error {
	if config.Previous == nil {
		return fmt.Errorf("you are on the first page")
	}

	locationsResp, err := config.Client.ListLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationsResp.Next
	config.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results{
		fmt.Println(loc.Name)
	}

	return nil
}


func commandExplore(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}

	name := args[0]

	fmt.Printf("Exploring %s...\n", name)

	locationArea, err := config.Client.ListLocationAreasPokemon(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters{
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
