package main

import (
	"fmt"
)

func commandExplore(conf *config) error {
	if len(conf.Args) == 0 {
		fmt.Println("Location argument not available for explore")
		return nil
	}
	response, err := conf.Client.GetPokemonByLocation(conf.Args[0])
	if err != nil {
		return fmt.Errorf("Error in getting response: %w", err)
	}
	for _, res := range response.PokemonEncounters {
		fmt.Printf("%s\n", res.Pokemn.Name)
	}
	return nil
}
