package main

import (
	"fmt"
	"math/rand"
)

type (
	kelvin float64
	sensor func() kelvin
)

func realSensor() kelvin {
	return 5
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	offset += 10
	fmt.Println(offset - 10)
	return func() kelvin {
		fmt.Println(offset)
		return s() + offset
	}
}

func main() {
	var (
		offset kelvin = 10
	)
	sensor1 := calibrate(realSensor, offset)
	sensor2 := calibrate(fakeSensor, offset)
	sensor3 := calibrate(fakeSensor, offset)
	fmt.Println(sensor1())
	fmt.Println(sensor2())
	fmt.Println(sensor3())
}
