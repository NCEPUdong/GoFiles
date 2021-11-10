package main

import "fmt"

type (
	celsius    float64
	fahrenheit float64
	kelvin     float64
)

// kelvinToCelsius converts °K to °C
func (k kelvin) kelvinToCelsius() kelvin {
	k = k - 273.15
	return k
}

// celsiusToFahrenheit converts celsius to fahrenheit
func (c celsius) celsiusToFahrenheit() celsius {
	c = c*9.0/5.0 + 32.0
	return c
}

// kelvinToFahrenheit converts kelvin to fahrenheit
func (k kelvin) kelvinToFahrenheit() kelvin {
	k = (k-273.15)*9.0/5.0 + 32.0
	return k
}

func main() {
	var (
		kelvin1 kelvin = 294.0
		kelvin2 kelvin = 0
	)
	fmt.Printf("%7.2f° K is %7.2f° C\n", kelvin1, kelvin1.kelvinToCelsius())
	fmt.Printf("%7.2f° K is %7.2f° K", kelvin2, kelvin1.kelvinToFahrenheit())
}
