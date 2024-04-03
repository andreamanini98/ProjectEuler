package main

import (
	"fmt"
	"time"
)

var totSundays int = 0

func main() {
	startDate := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)

	for !(startDate.Year() == 2000 && startDate.Month().String() == "December" && startDate.Day() == 31) {
		if startDate.Day() == 1 && startDate.Weekday().String() == "Sunday" {
			totSundays++
		}
		startDate = startDate.AddDate(0, 0, 1)
	}

	fmt.Println(totSundays)
}
