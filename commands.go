package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/W0ndermaker/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string, *Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func commandExit(locationArea string, config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(locationArea string, config *Config) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	for key, val := range getCommands() {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return nil
}

func commandMap(locationArea string, config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != nil {
		url = *config.Next
	}

	locatianAreas, err := config.pokeapiClient.GetLocationArea(url)
	if err != nil {
		return err
	}

	config.Next = locatianAreas.Next
	config.Previous = locatianAreas.Previous

	for _, location := range locatianAreas.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(locationArea string, config *Config) error {
	url := ""
	if config.Previous != nil {
		url = *config.Previous
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	locatianAreas, err := config.pokeapiClient.GetLocationArea(url)
	if err != nil {
		return err
	}

	config.Next = locatianAreas.Next
	config.Previous = locatianAreas.Previous

	for _, location := range locatianAreas.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(locationArea string, config *Config) error {
	if locationArea == "" {
		fmt.Println("Please provide a location area to explore")
		return nil
	}
	fmt.Printf("Exploring %v...\n", locationArea)
	url := "https://pokeapi.co/api/v2/location-area/" + locationArea

	locatianAreas, err := config.pokeapiClient.ExploreArea(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range locatianAreas.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil

}

func commandCatch(pokemonName string, config *Config) error {
	if pokemonName == "" {
		fmt.Println("Please provide a pokemon to catch")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	pokee, err := config.pokeapiClient.GetPokemon(url)
	if err != nil {
		return err
	}

	catchProbability := rand.Intn(pokee.BaseExperience * 2)
	if catchProbability > pokee.BaseExperience {
		fmt.Printf("%v was caught\n", pokee.Name)
		config.caughtPokemon[pokee.Name] = pokee
	} else {
		fmt.Printf("%v escaped\n", pokee.Name)
	}

	return nil
}

func commandInspect(pokemonName string, config *Config) error {
	if pokemon, ok := config.caughtPokemon[pokemonName]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Println("Name: ", pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Weight: ", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf(" -%v\n", t.Type.Name)
		}
	}
	return nil
}

func commandPokedex(s string, config *Config) error {
	fmt.Println("Your Pokedex")
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemons in the given location-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Your pokedex",
			callback:    commandPokedex,
		},
	}
}
