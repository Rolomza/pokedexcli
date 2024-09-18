package main

import (
	"bufio"
	"fmt"
	"os"
)

var cliCommandMap map[string]cliCommand

func init() {
	cliCommandMap = map[string]cliCommand{
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

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, command := range cliCommandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	return nil
}

func main() {

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		command, ok := cliCommandMap[scanner.Text()]
		if !ok {
			fmt.Println("Error: command not found. Type 'help' to see available commands.")
			continue
		}

		switch command.name {
		case "help":
			command.callback()
			continue
		case "exit":
			command.callback()
			return
		}
	}
	
}
