package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreasResp.Next
	cfg.previousLocationsURL = locationAreasResp.Previous

	for _, locationArea := range locationAreasResp.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreasResp.Next
	cfg.previousLocationsURL = locationAreasResp.Previous

	for _, locationArea := range locationAreasResp.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
