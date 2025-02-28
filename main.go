package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/bsuvonov/pokedex/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func( *Config, ...string) error
}


func getCommands() map[string]cliCommand {

	commands := map[string]cliCommand {
		"exit": {
			name:			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name:			"help",
			description: 	"Displays a help message",
			callback: 		commandHelp,
		},
		"map": {
			name:			"map",
			description: 	"Displays next 20 locations",
			callback: 		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays previous 20 locations",
			callback: 		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Explores pokemon in the location",
			callback:		commandExplore,
		},
		"catch": {
			name: 			"catch",
			description:	"Tries catching the pokemon by throwing a pokeball at it",
			callback: 		commandCatch,
		},
		"inspect": {
			name: 			"inspect",
			description: 	"Gets details of the caught pokemon",
			callback: 		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description: 	"Lists names of all the caught pokemon in Pokedex",
			callback:		commandPokedex,
		},
	}

	return commands
}

func main() {

	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	conf := Config{Previous: "", Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"}
	conf.Cache = pokecache.NewCache(5000*time.Millisecond)
	conf.Pokedex = make(map[string]Pokemon)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		fields := strings.Fields(strings.ToLower(input))
		command, exists := commands[fields[0]]

		if exists {
			err := command.callback(&conf, fields[1:]...)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}



func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}