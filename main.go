package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/RussNavas/pokedex/internal/pokeapi"
)


func main() {

	// make client
	client := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	// make config
	config := &Config{
		Client: &client,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		
		commandName := words[0]

		args := []string{}
		if len(words) > 1{
			args = words[1:]
		}


		command, exists := commands[commandName]
		if exists{
			err := command.callback(config, args)
			if err != nil{
				fmt.Println(err)
			}
		}else{
			fmt.Println("Unknown command")
		}
	}
}
