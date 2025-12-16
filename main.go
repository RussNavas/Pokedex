package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {

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
			err := command.callback()
			if err != nil{
				fmt.Println(err)
			}
		}else{
			fmt.Println("Unknown command")
		}
	}
}
