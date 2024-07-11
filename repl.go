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

		availableCommands := getCommands()

		commandName := cleanedPrompt[0]
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Command not supported")
			continue
		}

		command.callback()
		//fmt.Println("echoing...", cleanedPrompt)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(userText string) []string {
	lower := strings.ToLower(userText)
	words := strings.Fields(lower)
	return words
}
