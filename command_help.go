package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}

	var cmdOutput string
	for _, cmd := range getSupportedCommands() {
		cmdOutput += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n%s\n", cmdOutput)
	return nil
}
