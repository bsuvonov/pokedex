package main

import (
	"errors"
	"fmt"
)





func commandPokedex (conf *Config, args ...string) error {
	if len(conf.Pokedex) == 0{
		return errors.New("Your pokedex is empty!\n")
	} else {
		fmt.Println("Your Pokedex:")
		for pokemon_name, _ := range conf.Pokedex {
			fmt.Println(" -", pokemon_name)
		}
	}
	return nil
}