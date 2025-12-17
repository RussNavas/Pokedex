package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/RussNavas/pokedex/internal/pokeapi"
)


func main() {

	client := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	config := &Config{
		Client: &client,
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

		command, exists := commands[commandName]
		if exists{
			err := command.callback(config)
			if err != nil{
				fmt.Println(err)
			}
		}else{
			fmt.Println("Unknown command")
		}
	}
}
