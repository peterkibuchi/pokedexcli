package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
