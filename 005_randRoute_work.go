package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		spaceline = [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
		tripType  = [2]string{"Round-trip", "One-way"}
		trip      = "One-way"
		days      = 0
		price     = 0 // 3600-5000
	)
	fmt.Printf("%-20v      %-10v    %-10v      %-10v\n", "Spaceline", "Days", "Triptype", "Price")
	fmt.Println(strings.Repeat("=", 72))
	for i := 0; i < 10; i++ {
		trip = tripType[rand.Intn(2)]
		days = TripDays(trip)
		price = tripPrice(trip)
		fmt.Printf("%-20v      %-10v    %-10v      $ %-10v\n", spaceline[rand.Intn(3)], days, trip, price)
	}
}

func TripDays(trip string) int {
	speed := rand.Intn(14) + 17
	distance := 62100000
	days := int(distance/speed) + 1
	// return days
	if trip == "One-way" {
		return days
	} else {
		return days * 2
	}
}

func tripPrice(trip string) int {
	price := (rand.Intn(1400) + 3700) // 3600-5000
	if trip == "One-way" {
		return price
	} else {
		return price * 2
	}
}
