package main

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
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
	Pokedex  map[string]pokeapi.Pokemon
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
		"catch": {
		name:		"catch",
		description: "Try catching a pokemon listed in an area!",
		callback:	commandCatch,
		},
		"inspect": {
		name:		"inspect",
		description: "Use the Pokedex to inspect a pokemon that you have caught",
		callback:	commandInspect,
		},
		"pokedex": {
		name:		"pokedex",
		description: "List all captured Pokemon",
		callback:	commandPokedex,
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
		fmt.Println("")
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

func commandCatch(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide the name of a pokemon start capture attempt")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := config.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	const threshold = 40
	baseXP := pokemon.BaseExperience
	res := rand.Intn(baseXP)

	if res > threshold{
		fmt.Printf("%s escaped!\n", name)
	}

	fmt.Printf("%s was caught!\n", name)
	fmt.Printf("You may now inspect it with the inspect command.\n")
	config.Pokedex[name] = pokemon
	return nil
}

func commandInspect(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("provide the name of a caught pokemon to inspect")
	}

	name := args[0]

	// check pokedex
	if _, exists := config.Pokedex[name]; !exists{
		return fmt.Errorf("you have not caught that pokemon")
	}
	pokemon := config.Pokedex[name]
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats{
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, Type := range pokemon.Types{
		fmt.Printf("  - %v\n", Type.Type.Name)
	}
	return nil
}

func commandPokedex (config *Config, args []string) error {
	if len(config.Pokedex) == 0{
		return fmt.Errorf("you have 0 pokemon in your Pokedex")
	}else{
		fmt.Printf("Your Pokedex:\n")
		for _, pokemon := range config.Pokedex{
			fmt.Printf(" - %v\n", pokemon.Name)
		}
		return nil
	}
}
