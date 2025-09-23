package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandMap(conf *config) error {
	res, err := http.Get(conf.Next)
	if err != nil {
		return fmt.Errorf("Error in API: %w", err)
	}
	defer res.Body.Close()
	var locationMap *LocationMap
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response: %w", err)
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
	res, err := http.Get(conf.Previous)
	if err != nil {
		return fmt.Errorf("Error in API: %w", err)
	}
	defer res.Body.Close()
	var locationMap *LocationMap
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error in reading response: %w", err)
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
