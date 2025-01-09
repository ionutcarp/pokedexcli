package pokeapi

import (
	"github.com/ionutcarp/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(httpClientTimeout, cacheLifetimeInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheLifetimeInterval),
		httpClient: http.Client{
			Timeout: httpClientTimeout,
		},
	}
}
