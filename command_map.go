package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationReps, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationReps.Next
	cfg.prevLocationsURL = locationReps.Previous

	for _, loc := range locationReps.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
