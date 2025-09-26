package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandMap(conf *config) error {
	var data []byte
	var locationMap *LocationMap
	data, isCached := conf.CacheData.Get(conf.Next)
	if !isCached {
		res, err := http.Get(conf.Next)
		if err != nil {
			return fmt.Errorf("Error in API: %w", err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error in reading response: %w", err)
		}
		// Adding to cache
		conf.CacheData.Add(conf.Next, data)
	}
	if err := json.Unmarshal(data, &locationMap); err != nil {
		return fmt.Errorf("Error in Deserializing response: %w", err)
	}
	conf.Next = locationMap.Next
	conf.Previous = locationMap.Previous
	for _, resp := range locationMap.Results {
		fmt.Fprintln(os.Stdout, resp.Name)
	}
	return nil
}

func commandMapb(conf *config) error {
	if conf.Previous == "" {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	var data []byte
	data, isCached := conf.CacheData.Get(conf.Previous)
	var locationMap *LocationMap
	if !isCached {
		res, err := http.Get(conf.Previous)
		if err != nil {
			return fmt.Errorf("Error in API: %w", err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error in reading response: %w", err)
		}
		// caching data
		conf.CacheData.Add(conf.Previous, data)
	}
	if err := json.Unmarshal(data, &locationMap); err != nil {
		return fmt.Errorf("Error in Deserializing response: %w", err)
	}
	conf.Next = locationMap.Next
	conf.Previous = locationMap.Previous
	for _, resp := range locationMap.Results {
		fmt.Fprintln(os.Stdout, resp.Name)
	}
	return nil
}
