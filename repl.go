package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays the help menu",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: callbackExit,
		},
		"map": {
			name: "map",
			description: "Displays the next list of location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous list of location areas",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays the Pokemon in a location area.",
			callback: callbackExplore,
		},
	}
}