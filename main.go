package main

import (
	"time"

	"internal/pokeApi"
)

func main() {
	pokeClient := pokeApi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
