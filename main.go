package main

import (
	"fmt"
	"strconv"
	"ztm_scanner/cli/operations"
	"ztm_scanner/mappers"
	"ztm_scanner/providers/web"
	"ztm_scanner/ztm"
)

const (
	departuresLink = "https://ckan2.multimediagdansk.pl/departures?stopId="
	welcomeMsg     = `------------
Welcome to the ZTM Scanner.
This application is designed to display public transport schedules obtained from ZTM open data.
------------`
	menuMsg = `Please select one of the options to find a public transport stop.

[1] - Find by Stop ID
[2] - Find by Stop Name
[q] - Quit`
)

func main() {
	scheduleMapper := mappers.JsonMapper[ztm.Schedule]{}
	scheduleProvider := web.ResponseBodyProvider{}

	stopsMapper := mappers.JsonMapper[[]ztm.Stop]{}
	stopsProvider := web.StopsProvider{}
	userInput := ""

	fmt.Println(welcomeMsg)
mainLoop:
	for {
		fmt.Println(menuMsg)
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Printf("Input Error! %s\n", err)
			continue
		}
		switch userInput {
		case "1":
			fmt.Printf("Enter Stop ID: ")
			_, err := fmt.Scan(&userInput)
			if err != nil {
				fmt.Printf("Input Error! %s\n", err)
				continue
			}
		case "2":
			stopId, err := operations.FindStopsAndChoose(&stopsProvider, &stopsMapper)
			if err != nil {
				fmt.Printf("Error! %s\n", err)
				continue
			}
			if stopId == -1 {
				continue
			}
			userInput = strconv.Itoa(stopId)
		case "q":
			break mainLoop
		default:
			fmt.Printf("\nInvalid option.\n")
			continue
		}
		scheduleProvider.Url = departuresLink + userInput
		err = operations.GetAndDisplaySchedule(&scheduleProvider, &scheduleMapper)
		if err != nil {
			fmt.Printf("Error! %s\n", err)
		}
	}
}
