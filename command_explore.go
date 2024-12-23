package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("this command needs an area to explore")
	}

	name := args[0]
	fmt.Printf("Exploring %s...\n", name)
	locationDetailsResp, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetailsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
