package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		comm := ""
		sc.Scan()
		comm += sc.Text()
		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input", err)
		}
		commands := cleanInput(comm)
		firstWord := strings.ToLower(commands[0])
		if len(commands) > 0 {
			fmt.Printf("Your command was: %s\n", firstWord)
		}

	}

}

func cleanInput(text string) []string {
	result := []string{}
	cleanText := strings.TrimSpace(text)
	temp := ""
	emptySpace := ' '
	for _, letter := range cleanText {
		if letter == emptySpace {
			result = append(result, temp)
			temp = ""
		} else {
			temp += string(letter)
		}
	}
	result = append(result, temp)
	return result
}
