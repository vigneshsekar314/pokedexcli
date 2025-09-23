package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cliMapper := getCliMapper()
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
			if err := cmd.callback(); err != nil {
				fmt.Fprintln(os.Stderr, "Error in callback: ", err)
			}
		}
	}
}

func commandExit() error {
	fmt.Fprintln(os.Stdout, "Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Fprintln(os.Stdout, "Welcome to the Pokedex!\nUsage:")
	fmt.Fprintln(os.Stdout, "")
	cliMapper := getCliMapper()
	for _, cm := range cliMapper {
		fmt.Fprintf(os.Stdout, fmt.Sprintf("%s: %s\n", cm.name, cm.description))
	}
	return nil
}

func getCliMapper() map[string]cliCommand {
	return map[string]cliCommand{
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
