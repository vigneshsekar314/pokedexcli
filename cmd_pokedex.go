package main

import "fmt"

func commandPokedex(c *config) error {
	fmt.Printf("Your Pokedex:\n")
	for pokeName := range c.Pokeballs.Pokelist {
		fmt.Printf(" - %s\n", pokeName)
	}
	return nil
}
