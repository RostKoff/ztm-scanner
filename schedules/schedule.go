package schedules

import (
	"encoding/json"
	"fmt"
	"time"
	"ztm_scanner/web/client"
)

const (
	departureUrl = "https://ckan2.multimediagdansk.pl/departures?stopId="
)

type Departure struct {
	RouteId       int
	Headsign      string
	EstimatedTime time.Time
}

type Schedule struct {
	LastUpdate time.Time
	Departures []Departure
}

// Gets the JSON of the transport schedule for the stop identified by stopId and converts it to a Schedule struct.
func GetSchedule(stopId string) (Schedule, error) {
	s := Schedule{}
	respHolder, err := client.GetResponse(departureUrl + stopId)
	if err != nil {
		return s, fmt.Errorf("failed GetSchedule: %w", err)
	}
	if respHolder.StatusCode == 400 {
		return s, fmt.Errorf("stop with ID \"%s\" not found", stopId)
	}
	if !json.Valid(respHolder.Body) {
		return s, fmt.Errorf("valid JSON was not received")
	}
	err = json.Unmarshal(respHolder.Body, &s)
	if err != nil {
		err = fmt.Errorf("failed GetSchedule: %w", err)
	}
	return s, err
}
