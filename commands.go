package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getSupportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	var cmdOutput string
	for _, cmd := range getSupportedCommands() {
		cmdOutput += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n%s\n", cmdOutput)
	return nil
}