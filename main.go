package main

import (
	"fmt"
	"log"

	"github.com/vanitasko/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	// StartRepl()
}
