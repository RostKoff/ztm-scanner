package operations

import (
	"fmt"
	"strconv"
	"strings"
	"ztm_scanner/cli/display"
	"ztm_scanner/mappers"
	"ztm_scanner/ztm"
)

type mapper[T any] interface {
	// Retrieves data from the provider and stores it in the value pointed to by 't'.
	MapValue(p mappers.Provider, t *T) error
}

// Retrieves data on public transport schedule from the provider and then displays it in the console.
// It also gives the user the opportunity to retrieve the data again, potentially updating it.
func GetAndDisplaySchedule(p mappers.Provider, m mapper[ztm.Schedule]) error {
	schedule := ztm.Schedule{}
	userInput := ""
	invalidOption := false
retriveDate:
	for {
		err := m.MapValue(p, &schedule)
		if err != nil {
			return fmt.Errorf("failed GetAndDisplaySchedule: %w", err)
		}
		for {
			err = display.DisplaySchedule(schedule)
			if err != nil {
				return fmt.Errorf("failed GetAndDisplaySchedule: %w", err)
			}
			if invalidOption {
				fmt.Printf("Invalid option.\n")
				invalidOption = false
			}
			fmt.Printf("Choose an option\n[r] - Refresh Data\t[q] - Go back\n")

			_, err = fmt.Scan(&userInput)
			if err != nil {
				return fmt.Errorf("failed GetAndDisplaySchedule: %w", err)
			}
			switch userInput {
			case "r":
				continue retriveDate
			case "q":
				break retriveDate
			default:
				invalidOption = true
			}
		}
	}
	return nil
}

// Retrieves information about all public transport stops, asks the user to enter the name of the stop and displays the filtered results in the console.
// After that, the user can select one of the filtered stops, and the function will return stopId. If the user has not selected anything, the function will return -1.
// The function also gives the user the opportunity to retrieve the data again, potentially updating it, and change the name by which the search is performed.
func FindStopsAndChoose(p mappers.Provider, m mapper[[]ztm.Stop]) (int, error) {
	userInput := ""
	stopName := ""
	errorMsg := ""
	invalidOption := false
	changeName := true
retriveData:
	for {
		stops := make([]ztm.Stop, 0)
		err := m.MapValue(p, &stops)
		if err != nil {
			return 0, fmt.Errorf("failed FindStopsAndChoose: %w", err)
		}
	chooseName:
		for {
			if changeName {
				fmt.Printf("Enter Stop Name\n-For example 'Dworzec'-\n")
				_, err = fmt.Scan(&stopName)
				if err != nil {
					return 0, fmt.Errorf("failed FindStopsAndChoose: %w", err)
				}
				changeName = false
			}
			filteredStops := FilterStopsByName(stopName, stops)
			for {
				display.DisplayStops(filteredStops)
				if invalidOption {
					fmt.Printf("%s\n", errorMsg)
					invalidOption = false
				}
				fmt.Printf("Enter 'Nr' of the desired stop or one of the following options\n[r] - Refresh data\t[n] - Change name\t[q] - Go back\n")

				_, err = fmt.Scan(&userInput)
				if err != nil {
					return 0, fmt.Errorf("failed FindStopsAndChoose: %w", err)
				}
				switch userInput {
				case "r":
					continue retriveData
				case "n":
					changeName = true
					continue chooseName
				case "q":
					break retriveData
				default:
					num, err := strconv.Atoi(userInput)
					if err != nil || num < 0 || num >= len(filteredStops) {
						errorMsg = "Error! Entered invalid value."
						invalidOption = true
						continue
					}
					return filteredStops[num].StopId, nil
				}
			}
		}
	}
	return -1, nil
}

func FilterStopsByName(name string, stops []ztm.Stop) []ztm.Stop {
	filteredStops := make([]ztm.Stop, 0)

	for _, s := range stops {
		if strings.Contains(strings.ToLower(s.StopName), strings.ToLower(name)) {
			filteredStops = append(filteredStops, s)
		}
	}
	return filteredStops
}
