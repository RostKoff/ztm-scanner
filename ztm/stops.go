package ztm

import (
	"strings"
)

type Stop struct {
	StopId   int
	StopCode string
	StopName string
}

func FilterStopsByName(name string, stops []Stop) []Stop {
	filteredStops := make([]Stop, 0)

	for _, s := range stops {
		if strings.Contains(strings.ToLower(s.StopName), strings.ToLower(name)) {
			filteredStops = append(filteredStops, s)
		}
	}
	return filteredStops
}
