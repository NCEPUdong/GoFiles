package main

import "fmt"

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// celsiusToFahrenheit converts celsius to fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	c = c*9.0/5.0 + 32.0
	return c
}

// kelvinToFahrenheit converts kelvin to fahrenheit
func kelvinToFahrenheit(k float64) float64 {
	k = kelvinToCelsius(k)
	k = celsiusToFahrenheit(k)
	return k
}

func main() {
	var (
		kelvin  float64 = 294.0
		kelvin0 float64 = 0
	)
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%7.2f° K is %7.2f° C\n", kelvin, celsius)
	fmt.Printf("%7.2f° K is %7.2f° C", kelvin0, kelvinToFahrenheit(float64(kelvin0)))
}
