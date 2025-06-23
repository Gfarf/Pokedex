package pokeApi

import (
	"net/http"
	"time"

	"internal/pokecache"
)

// Client -
type Client struct {
	Pokedex    map[string]RespPokStats
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Pokedex: make(map[string]RespPokStats),
		cache:   pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
