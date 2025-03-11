# Pokedex
## Description
A pokedex simulation CLI app. This project uses is mainly a Go practice with APIs, and uses [PokéAPI](https://pokeapi.co/).

## How to Use This Project
Run commands below to download the project:
```bash
$ git clone https://github.com/bsuvonov/pokedex
$ cd pokedex
```
Make sure to install [Go](https://go.dev/). Then you can build the project or run it using `go run .`.
How to build the project:
```bash
$ go build
$ ./pokedex
```
## How to Play The Game
Here is the list of commands:

| Command                   | Description                                       |
| ------------------------- | --------------------------------------------------|
| `exit`                    | Exit the Pokedex                                  |
| `help`                    | Displays a help message                           |
| `map`                     | Displays next 20 locations in Pokemon World       |
| `mapb`                    | Displays previous 20 locations in Pokemon World   |
| `explore <location-area>` | Explores pokemon in the location                  |
| `catch <pokemon>`         | Tries catching the pokemon by throwing a pokeball |
| `inspect <pokemon>`       | Gets details of the caught pokemon                |
| `pokedex`                 | Lists names of all the caught pokemon in Pokedex  |
