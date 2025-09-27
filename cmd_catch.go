package main

import (
	"fmt"
	"github.com/vigneshsekar314/pokedexcli/internal/pokeapi"
	"math/rand"
)

func commandCatch(conf *config) error {
	if len(conf.Args) == 0 {
		fmt.Println("Pokemon name should be provided")
		return nil
	}
	pokeName := conf.Args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	response, err := conf.PokeClient.GetPokemonInfo(pokeName)
	if err != nil {
		return fmt.Errorf("Error in getting response: %w", err)
	}
	prob := rand.Intn(500) // base experience is below 500 for most pokemons viewed
	threshold := min(response.BaseExperience, 450)
	if prob > threshold {
		fmt.Printf("%s was caught!\n", pokeName)
		if conf.Pokeballs.Pokelist == nil {
			conf.Pokeballs.Pokelist = map[string]pokeapi.PokemonInfo{}
		}
		_, ok := conf.Pokeballs.Pokelist[pokeName]
		if !ok {
			conf.Pokeballs.Pokelist[pokeName] = response
		}
		fmt.Printf("%s info: name=%s, base experience=%d, id=%d\n", pokeName, response.Name, response.BaseExperience, response.Id)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokeName)
	return nil
}
