package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)


type EncountersAtLocation struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}



func commandExplore(conf *Config, args ...string) error {

	base_url := "https://pokeapi.co/api/v2/location-area/"
	for _, arg := range args {
		fmt.Println("Exploring", arg + "...")

		res, err := http.Get(base_url + arg)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return errors.New("Response failed with status code: " + strconv.Itoa(res.StatusCode) + " and\nbody: " + string(body) + "\n")
		}
		if err!= nil {
			return err
		}
		
		var encounters EncountersAtLocation
		
		if err := json.Unmarshal(body, &encounters); err != nil {
			return err
		}

		if len(encounters.PokemonEncounters) > 0 {
			fmt.Println("Found Pokemon:")
		}
		for _, pokemon := range encounters.PokemonEncounters {
			fmt.Println(" - " + pokemon.Pokemon.Name)
		}
	}

	return nil
}




