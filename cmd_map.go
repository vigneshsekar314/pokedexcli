package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandMap(conf *config) error {
	fmt.Printf("url is: %s\n", conf.Next)
	res, err := http.Get(conf.Next)
	if err != nil {
		// fmt.Printf("Error in reading response: %s", err)
		return fmt.Errorf("Error in API: %w", err)
	}
	defer res.Body.Close()
	var locationMap LocationMap
	data, err := io.ReadAll(res.Body)
	if err != nil {
		// fmt.Printf("Error in reading response: %s", err)
		return fmt.Errorf("Error in reading response: %w", err)
	}
	if err := json.Unmarshal(data, &locationMap); err != nil {
		// fmt.Printf("Error in reading response: %s", err)
		return fmt.Errorf("Error in Deserializing response: %w", err)
	}
	locRs, err := json.Marshal(locationMap)
	if err != nil {
		// fmt.Printf("Error in reading response: %s", err)
		return fmt.Errorf("error: %w", err)
	}
	conf.Next = locationMap.next
	conf.Previous = locationMap.previous
	fmt.Printf("response serialized: %s\n", locRs)
	fmt.Printf("response: %v and url: %v\n", locationMap.results, conf.Next)
	for _, resp := range locationMap.results {
		fmt.Fprintln(os.Stdout, resp.name)
	}
	return nil
}

type LocationMap struct {
	count    int
	next     string
	previous string
	results  []LocationNmUrl
}

type LocationNmUrl struct {
	name string
	url  string
}
