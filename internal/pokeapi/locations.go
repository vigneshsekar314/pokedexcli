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
	// var data []byte
	// var locationMap LocationMap
	// data, isCached := c.cacheData.Get(url)
	// if !isCached {
	// 	req, err := http.NewRequest("GET", url, nil)
	// 	if err != nil {
	// 		return LocationMap{}, fmt.Errorf("Error in creating a new request: %w", err)
	// 	}
	// 	res, err := c.httpClient.Do(req)
	// 	if err != nil {
	// 		return LocationMap{}, fmt.Errorf("Error in API: %w", err)
	// 	}
	// 	defer res.Body.Close()
	// 	data, err = io.ReadAll(res.Body)
	// 	if err != nil {
	// 		return LocationMap{}, fmt.Errorf("Error in reading response: %w", err)
	// 	}
	// 	// Adding to cache
	// 	c.cacheData.Add(url, data)
	// }
	// if err := json.Unmarshal(data, &locationMap); err != nil {
	// 	return LocationMap{}, fmt.Errorf("Error in Deserializing response: %w", err)
	// }
	// return locationMap, nil
}

func (c *Client) GetPokemonByLocation(location string) (PokemonsByLocation, error) {
	url := BaseUrl + "location-area/" + location
	resp, err := GetResponse[PokemonsByLocation](c, url, "Invalid location")
	if err != nil {
		return PokemonsByLocation{}, err
	}
	//data, isCached := c.cacheData.Get(url)
	//if !isCached {
	//	req, err := http.NewRequest("GET", url, nil)
	//	if err != nil {
	//		return PokemonsByLocation{}, fmt.Errorf("Error in creating new request: %w", err)
	//	}
	//	res, err := c.httpClient.Do(req)
	//	if err != nil {
	//		return PokemonsByLocation{}, fmt.Errorf("Error in retrieving response: %w", err)
	//	}
	//	defer res.Body.Close()
	//	if res.StatusCode != 200 {
	//		fmt.Printf("Invalid location\n")
	//		return PokemonsByLocation{}, nil
	//	}
	//	data, err = io.ReadAll(res.Body)
	//	if err != nil {
	//		return PokemonsByLocation{}, fmt.Errorf("Error in reading response bytes: %w", err)
	//	}
	//	//Adding to cache
	//	c.cacheData.Add(url, data)
	//}
	//if err := json.Unmarshal(data, &pokemonsByLocation); err != nil {
	//	return PokemonsByLocation{}, fmt.Errorf("Error in unmarshalling: %w", err)
	//}
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
