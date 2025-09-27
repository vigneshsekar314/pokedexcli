package main

import (
	"fmt"
)

func commandInspect(conf *config) error {
	if len(conf.Args) == 0 {
		return fmt.Errorf("Pokemon name is required")
	}
	pokeName := conf.Args[0]
	pokeInfo, ok := conf.Pokeballs.Pokelist[pokeName]
	if !ok {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n", pokeInfo.Name)
	fmt.Printf("Height: %d\n", pokeInfo.Height)
	fmt.Printf("Weight: %d\n", pokeInfo.Weight)
	fmt.Printf("Stats:\n")
	for _, item := range pokeInfo.StatList {
		fmt.Printf("  -%s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, item := range pokeInfo.Types {
		fmt.Printf("  - %s\n", item.Type.Name)
	}
	return nil
}
