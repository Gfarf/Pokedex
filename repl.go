package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		t := scanner.Text()
		inputs := cleanInput(t)
		j := getCommands()
		for command := range j {
			if inputs[0] == command {
				j[command].callback()
			}
		}

	}
}
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)
	return res
}
