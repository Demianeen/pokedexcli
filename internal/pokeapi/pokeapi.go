package pokeapi

import (
	"net/http"
	"time"

	"github.com/Demianeen/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(cacheDuration time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheDuration),
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
