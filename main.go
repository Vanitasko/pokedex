package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		//Grab new prompt
		fmt.Print("pokeDEX => ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		prompt := scanner.Text()
		fmt.Println("echoing...", prompt)
	}
}
