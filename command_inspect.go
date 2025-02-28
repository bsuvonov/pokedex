package main

import (
	"errors"
	"fmt"
)





func commandInspect(conf *Config, args ...string) error {
	for _, pokemon_name := range args {
		pokemon, ok := conf.Pokedex[pokemon_name]
		if !ok {
			return errors.New("you have not caught that pokemon\n")
		}
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Println("  -" + stat.Stat.Name + ":", stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typee := range pokemon.Types {
			fmt.Println("  -", typee.Type.Name)
		}
	}
	return nil
}