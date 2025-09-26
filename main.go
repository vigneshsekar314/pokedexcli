package main

import (
	"bufio"
	"fmt"
	"github.com/vigneshsekar314/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cliMapper := getCliMapper()
	newCacheData := pokecache.NewCache(time.Second * 5)
	conf := config{
		Next:      "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous:  "",
		CacheData: &newCacheData,
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

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next      string
	Previous  string
	CacheData *pokecache.Cache
}

func getCliMapper() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays names of next 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
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
