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
	cliMapper := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provides help for the Pokedex",
			callback:    commandHelp,
		},
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
			if err := cmd.callback(); err != nil {
				fmt.Fprintln(os.Stderr, "Error in callback: ", err)
			}
			// switch command {
			// case "exit":
			// 	cmd, ok := cliMapper["exit"]
			// 	if !ok {
			// 		fmt.Printf("exit command not found in mapper")
			// 	}
			// 	if cmd.callback != nil {
			// 		if err := cmd.callback(); err != nil {
			// 			fmt.Fprintln(os.Stderr, "Error in callback command:", err)
			// 		}
			// 	}
			// default:
			// 	fmt.Fprintln(os.Stderr, "Unknown command")
			// }
		}
	}
}

func commandExit() error {
	fmt.Fprintln(os.Stdout, "Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Fprintln(os.Stdout, "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	return strings.Fields(result)
}
