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

		cleanedPrompt := cleanInput(prompt)
		if len(cleanedPrompt) == 0 {
			continue
		}

		command := cleanedPrompt[0]
		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Command not configured")
		}

		//fmt.Println("echoing...", cleanedPrompt)

	}
}

func cleanInput(userText string) []string {
	lower := strings.ToLower(userText)
	words := strings.Fields(lower)
	return words
}
