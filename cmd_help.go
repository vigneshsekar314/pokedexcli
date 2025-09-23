package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Fprintln(os.Stdout, "Welcome to the Pokedex!\nUsage:")
	fmt.Fprintln(os.Stdout, "")
	cliMapper := getCliMapper()
	for _, cm := range cliMapper {
		fmt.Fprintf(os.Stdout, fmt.Sprintf("%s: %s\n", cm.name, cm.description))
	}
	return nil
}
