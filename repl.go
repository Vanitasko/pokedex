package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	for {
		//Grab new prompt
		fmt.Print("pokeDEX => ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		prompt := scanner.Text()
		fmt.Println("echoing...", prompt)
	}
}

func cleanInput(userText string) []string {
	lower := strings.ToLower(userText)
	words := strings.Fields(lower)
	return words
}
