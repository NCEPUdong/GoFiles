package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		year        = 2018
		month       = 2
		daysInMonth = 31
	)
	switch month {
	case 2:
		if year%4 == 0 || year%100 != 0 {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	default:
		daysInMonth = 31
	}

	day := 0
	for i := 1; i <= 10; i++ {
		year = rand.Intn(2038-1970) + 1970
		day = rand.Intn(daysInMonth) + 1
		month = rand.Intn(12) + 1
		fmt.Println(era, year, month, day)
	}
}
