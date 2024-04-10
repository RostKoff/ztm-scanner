package display

import (
	"fmt"
	"time"
	"ztm_scanner/ztm"
)

func DisplaySchedule(schedule ztm.Schedule) error {
	loc, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		return fmt.Errorf("DisplaySchedule: %w", err)
	}

	printLine()
	fmt.Printf("Last Update: %s\n", schedule.LastUpdate.In(loc).Format(time.DateTime))
	printLine()
	fmt.Printf("Route\tHeadsign\tDeparture Time\n")
	if len(schedule.Departures) == 0 {
		fmt.Printf("-No routes found-\n")
	} else {
		for _, dep := range schedule.Departures {
			fmt.Printf("%d\t%s\t%s\n", dep.RouteId, dep.Headsign, dep.EstimatedTime.In(loc).Format(time.TimeOnly))
		}
	}
	printLine()

	return nil
}

func DisplayStops(stops []ztm.Stop) {
	printLine()
	fmt.Printf("Nr\tStop ID\tStop Name\n")
	if len(stops) == 0 {
		fmt.Printf("-No stops found-\n")
	} else {
		for n, s := range stops {
			fmt.Printf("%d\t%d\t%s %s\n", n, s.StopId, s.StopName, s.StopCode)
		}
	}
	printLine()
}

func printLine() {
	fmt.Printf("--------------------\n")
}
