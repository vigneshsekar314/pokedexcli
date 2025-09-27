package main

import (
	"fmt"
	"os"
)

func commandMapf(conf *config) error {
	locationMap, err := conf.Client.GetLocationList(conf.Next)
	if err != nil {
		return fmt.Errorf("Error in calling LocationList: %w\n", err)
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
	locationMap, err := conf.Client.GetLocationList(conf.Previous)
	if err != nil {
		return fmt.Errorf("Error in calling LocationList: %w\n", err)
	}
	conf.Next = locationMap.Next
	conf.Previous = locationMap.Previous
	for _, resp := range locationMap.Results {
		fmt.Fprintln(os.Stdout, resp.Name)
	}
	return nil
}
