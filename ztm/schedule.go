package ztm

import (
	"time"
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
