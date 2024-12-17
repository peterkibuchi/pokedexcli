package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	var cmdOutput string
	for _, cmd := range getSupportedCommands() {
		cmdOutput += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n%s\n", cmdOutput)
	return nil
}
