package main

import "fmt"

func main() {
	const (
		distance = 56000000
		days     = 28
		speed    = distance / days
	)
	fmt.Println("The speed towards Malacandra will be", speed, "km/s")
}
