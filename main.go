package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/W0ndermaker/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := Config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}
		command := words[0]
		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}
		if val, ok := getCommands()[command]; ok {
			val.callback(arg, &conf)
		} else {
			fmt.Println("Unknown command")
		}

	}
}
