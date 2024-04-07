package main

import (
	"fmt"
	"ztm_scanner/ztm"
)

func main() {
	res, err := ztm.GetStopsByName("bajana")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
