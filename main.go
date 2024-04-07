package main

import (
	"fmt"
	"ztm_scanner/schedules"
)

func main() {
	res, err := schedules.GetSchedule("2082")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
