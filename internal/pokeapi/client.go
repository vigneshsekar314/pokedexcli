package pokeapi

import (
	"github.com/vigneshsekar314/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	cacheData  pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cacheData: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
