package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("this command needs a pokemon to inspect")
	} else if len(args) > 1 {
		return errors.New("too many arguments")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Wright: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" - %v\n", typeInfo.Type.Name)
	}

	return nil
}
