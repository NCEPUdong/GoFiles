package main

import (
	"fmt"
)

type (
	tempConvert func(temp float64) float64
	table       func(int, float64, tempConvert)
)

func main() {
	var (
		temp1 float64
		// temp2 float64
	)
	drawBorder()
	drawHeader("ctf")
	drawBorder()
	for i := 0; i < 28; i++ {
		temp1 = -40.0 + float64(i)*5.0
		drawTable(float64(temp1), celsiusToFahrenheit)
	}
	drawBorder()
}

func drawTable(temp float64, con tempConvert) {
	fmt.Printf("| %8.2f | %8.2f |\n", temp, con(temp))
}

func drawHeader(style string) {
	if style == "ctf" {
		fmt.Printf("| %8s | %8s |\n", "°C", "°F")
	} else {
		fmt.Printf("| %8s | %8s |\n", "°F", "°C")
	}
}

func drawBorder() {
	fmt.Printf(func() string {
		var stuff string
		for i := 0; i < 23; i++ {
			stuff += "="
		}
		stuff += "\n"
		return stuff
	}())
}

func celsiusToFahrenheit(c float64) float64 {
	return float64(c*9.0/5.0 + 32.0)
}

func fahrenheitToCelsius(f float64) float64 {
	return float64((f - 32.0) * 5.0 / 9.0)
}
