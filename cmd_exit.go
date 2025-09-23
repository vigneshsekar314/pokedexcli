package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config) error {
	fmt.Fprintln(os.Stdout, "Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
