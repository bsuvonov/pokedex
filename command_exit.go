package main

import (
	"errors"
	"fmt"
	"os"
)


func commandExit(conf *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("closingggg")
}


