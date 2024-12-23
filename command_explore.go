package main

import (
	"fmt"
)

func commandExplore(cfg *config, param *string) error {
	fmt.Printf("Exploring %s...\n", *param)
	locationDetailsResp, err := cfg.pokeapiClient.ListLocationDetails(param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationDetailsResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
