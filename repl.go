package main

import(
	"strings"
	"fmt"
	"os"
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
	callback	func() error
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
	}
}


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands = getCommands()
	for name, command := range commands{
		fmt.Printf("%v: %v\n", name, command.description)
	}
	return nil
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
