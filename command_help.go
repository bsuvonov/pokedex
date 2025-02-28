package main

import (
	"fmt"
)


func commandHelp(conf *Config, args ...string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()

	for command, value := range commands {
		fmt.Println(command + ":", value.description)
	}
	return nil
}