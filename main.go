package main

import (
	"bufio"
	"fmt"
	"github.com/vigneshsekar314/pokedexcli/internal/pokeapi"
	"os"
	"strings"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cliMapper := getCliMapper()
	default_url := pokeapi.BaseUrl + "location-area/?offset=0&limit=20"
	default_prev_url := ""
	conf := config{
		Next:       default_url,
		Previous:   default_prev_url,
		PokeClient: pokeapi.NewClient(time.Second*60, time.Second*20),
	}
	for {
		fmt.Print("pokedex > ")
		sc.Scan()
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input", err)
		}
		commands := cleanInput(sc.Text())
		if len(commands) > 0 {
			command := commands[0]
			args := commands[1:]
			conf.Args = args
			cmd, ok := cliMapper[command]
			if !ok {
				fmt.Fprintln(os.Stderr, "Unknown command")
				continue
			}
			if cmd.callback == nil {
				fmt.Fprintln(os.Stderr, "Callback not defined")
				continue
			}
			if err := cmd.callback(&conf); err != nil {
				fmt.Fprintln(os.Stderr, "Error in callback: ", err)
			}
		}
	}
}

func getCliMapper() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays names of next 20 location areas in the Pokemon world.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays names of pokemon found in the location given.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon with the given name.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check information about a pokemon",
			callback:    commandInspect,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	return strings.Fields(result)
}
