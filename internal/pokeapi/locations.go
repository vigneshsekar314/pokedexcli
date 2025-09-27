package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(url string) (LocationMap, error) {

	var data []byte
	var locationMap LocationMap
	data, isCached := c.cacheData.Get(url)
	if !isCached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationMap{}, fmt.Errorf("Error in creating a new request: %w", err)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationMap{}, fmt.Errorf("Error in API: %w", err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationMap{}, fmt.Errorf("Error in reading response: %w", err)
		}
		// Adding to cache
		c.cacheData.Add(url, data)
	}
	if err := json.Unmarshal(data, &locationMap); err != nil {
		return LocationMap{}, fmt.Errorf("Error in Deserializing response: %w", err)
	}
	return locationMap, nil
}

func (c *Client) GetPokemonByLocation(location string) (PokemonsByLocation, error) {
	url := BaseUrl + "location-area/" + location
	data, isCached := c.cacheData.Get(url)
	var pokemonsByLocation PokemonsByLocation
	if !isCached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonsByLocation{}, fmt.Errorf("Error in creating new request: %w", err)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonsByLocation{}, fmt.Errorf("Error in retrieving response: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			fmt.Printf("Invalid location\n")
			return PokemonsByLocation{}, nil
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonsByLocation{}, fmt.Errorf("Error in reading response bytes: %w", err)
		}
		//Adding to cache
		c.cacheData.Add(url, data)
	}
	if err := json.Unmarshal(data, &pokemonsByLocation); err != nil {
		return PokemonsByLocation{}, fmt.Errorf("Error in unmarshalling: %w", err)
	}
	return pokemonsByLocation, nil
}
