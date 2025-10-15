package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		input = scanner.Text()
		result := cleanInput(input)
		if len(result) > 0 {
			fmt.Printf("Your command was: %s\n", result[0])
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	result := strings.Fields(lowered)
	return result
}
