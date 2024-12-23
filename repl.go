package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/peterkibuchi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		if err := reader.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		supportedCommands := getSupportedCommands()
		if command, exists := supportedCommands[commandName]; exists {
			if command.name == "explore" {
				if len(words) < 2 {
					fmt.Println("this command needs an area to explore")
					continue
				}
				param := words[1]
				err := command.callback(cfg, &param)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(cfg, nil)
				if err != nil {
					fmt.Println(err)
				}
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
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
		"explore": {
			name:        "explore",
			description: "Displays list of all the Pokémon located in a particular location area.",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previously displayed 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
