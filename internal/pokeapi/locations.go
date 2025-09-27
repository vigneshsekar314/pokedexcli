package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(url string) (LocationMap, error) {
	resp, err := GetResponse[LocationMap](c, url, "Invalid url")
	if err != nil {
		return LocationMap{}, err
	}
	return resp, nil
}

func (c *Client) GetPokemonByLocation(location string) (PokemonsByLocation, error) {
	url := BaseUrl + "location-area/" + location
	resp, err := GetResponse[PokemonsByLocation](c, url, "Invalid location")
	if err != nil {
		return PokemonsByLocation{}, err
	}
	return resp, nil
}

func GetResponse[T any](c *Client, url, errormsg string) (T, error) {
	var response T
	data, isCached := c.cacheData.Get(url)
	if isCached {
		if err := json.Unmarshal(data, &response); err != nil {
			return response, fmt.Errorf("Error in unmarshalling: %w", err)
		}
		return response, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, fmt.Errorf("Error in creating new request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("Error in retrieving response: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("%s\n", errormsg)
		return response, nil
	}
	data, err = BytesToJSON(res.Body, &response)
	if err != nil {
		return response, err
	}
	c.cacheData.Add(url, data)
	return response, nil
}

func BytesToJSON[T any](b io.ReadCloser, res *T) ([]byte, error) {
	data, err := io.ReadAll(b)
	if err != nil {
		return nil, fmt.Errorf("Error in reading bytes: %w", err)
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("Error in unmarshalling: %w", err)
	}
	return data, nil
}
