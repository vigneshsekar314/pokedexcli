package main

import (
	"github.com/vigneshsekar314/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
	Client   pokeapi.Client
	Args     []string
}
