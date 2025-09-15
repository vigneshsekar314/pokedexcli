package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	result := []string{}
	cleanText := strings.TrimSpace(text)
	temp := ""
	emptySpace := ' '
	for _, letter := range cleanText {
		fmt.Printf("character is - %c\n", letter)
		if letter == emptySpace {
			result = append(result, temp)
			fmt.Printf("word added: %s\n", temp)
			temp = ""
		} else {
			temp += string(letter)
		}
	}
	result = append(result, temp)
	return result
}
