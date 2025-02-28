package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"math/rand"
	"encoding/json"
)



type Pokemon struct {
	Name	string `json:"name"`
	Height	int `json:"height"`
	Weight	int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}




func commandCatch (conf *Config, args ...string) error{

	base_url := "https://pokeapi.co/api/v2/pokemon/"
	for _, pokemon_name := range args {

		res, err := http.Get(base_url + pokemon_name)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			if res.StatusCode == 404{
				return fmt.Errorf(fmt.Sprintf("Pokemon with the name %s was not found\n", pokemon_name))
			}
			return errors.New("Response failed with status code: " + strconv.Itoa(res.StatusCode) + " and\nbody: " + string(body) + "\n")
		}
		if err!= nil {
			return err
		}



		fmt.Println("Throwing a Pokeball at", pokemon_name + "...")

		rand_n := rand.Float64()

		if rand_n > 0.5 {
			fmt.Println(pokemon_name, "was caught!")

			var pokemon Pokemon
		
			if err := json.Unmarshal(body, &pokemon); err != nil {
				return err
			}
			conf.Pokedex[pokemon_name] = pokemon
		} else {
			fmt.Println(pokemon_name, "escaped!")
		}
	}

	return nil
}