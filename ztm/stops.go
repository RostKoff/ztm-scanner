package ztm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"ztm_scanner/web/client"
)

const (
	stopsUrl = "https://ckan.multimediagdansk.pl/dataset/c24aa637-3619-4dc2-a171-a23eec8f2172/resource/4c4025f0-01bf-41f7-a39f-d156d201b82b/download/stops.json"
)

type Stop struct {
	StopId   int
	StopCode string
	StopName string
}

// Gets the most up-to-date list of all stops in Gdańsk
func getStops() ([]Stop, error) {
	stops := make([]Stop, 0)

	resp, err := client.GetResponse(stopsUrl)
	if err != nil {
		return stops, fmt.Errorf("failed GetStopsByName: %w", err)
	}

	r := bytes.NewReader(resp.Body)
	dec := json.NewDecoder(r)

	// Skip unnecessary JSON part
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return stops, fmt.Errorf("failed GetStopsByName: %w", err)
		}
		if t == "stops" {
			break
		}
	}

	for dec.More() {
		err := dec.Decode(&stops)
		if err != nil {
			return stops, fmt.Errorf("failed GetStopsByName: %w", err)
		}
	}
	fmt.Println(len(stops))
	return stops, nil
}

func FilterStopsByName(name string, stops []Stop) []Stop {
	filteredStops := make([]Stop, 0)

	for _, s := range stops {
		if strings.EqualFold(name, s.StopName) {
			filteredStops = append(filteredStops, s)
		}
	}
	return filteredStops
}

func GetStopsByName(name string) ([]Stop, error) {
	allStops, err := getStops()
	if err != nil {
		return allStops, fmt.Errorf("failed GetStopsByName: %w", err)
	}

	return FilterStopsByName(name, allStops), nil
}
