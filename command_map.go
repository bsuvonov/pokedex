package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bsuvonov/pokedex/internal/pokecache"
)



type Config struct{
	Cache pokecache.Cache
	Next string
	Previous string
	Pokedex map[string]Pokemon
}

type Location struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(conf *Config, args ...string) error{

	if conf.Next!=""{
		url := conf.Next

		body, ok := conf.Cache.Get(url)
		if !ok {
		

			res, err := http.Get(url)
			if err != nil {
				return err
			}
			body, err = io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				return errors.New("Response failed with status code: " + strconv.Itoa(res.StatusCode) + " and\nbody: " + string(body) + "\n")
			}
			if err != nil {
				return err
			}
		}
		var location Location
		if err := json.Unmarshal(body, &location); err != nil {
			return err
		}

		for _, result := range location.Results {
			fmt.Println(result.Name)
		}

		conf.Cache.Add(url, body)
		conf.Previous = url
		conf.Next = location.Next
	} else {
		return errors.New("reached the end of locations")
	}
	return nil
}


func commandMapb(conf *Config, args ...string) error{
	if conf.Previous != "" {
		url := conf.Previous

		res, err := http.Get(url)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return errors.New("Response failed with status code: " + strconv.Itoa(res.StatusCode) + " and\nbody: " + string(body) + "\n")
		}
		if err != nil {
			return err
		}
		var location Location
		if err := json.Unmarshal(body, &location); err != nil {
			return err
		}

		for _, result := range location.Results {
			fmt.Println(result.Name)
		}

		conf.Cache.Add(url, body)
		conf.Next = url
		if location.Previous != nil {
			conf.Previous = location.Previous.(string)
		} else {
			conf.Previous = ""
		}
	} else {
		fmt.Println("you're on the first page")
	}
	return nil
}