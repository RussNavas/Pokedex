package main

import(
	"strings"
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