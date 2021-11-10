package main

import (
	"fmt"
	"math"
)

type (
	location struct {
		name      string
		lat, long float64
	}

	world struct {
		radius float64
	}

	// (Global Positioning System, GPS)
	gps struct {
		location1, location2 location
		world
	}

	rover struct {
		gps
	}

	// distancer interface {
	// 	distance() float64
	// }
)

// rad converts degrees to radius
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (l location) description() string {
	return fmt.Sprintf("name: %s, latitude: %f, longitude: %f", l.name, l.lat, l.long)
}

// distance calculation using the Spherical Law of Cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) distance(p1, p2 location) float64 {
	return g.world.distance(g.location1, g.location2)
}

func (g gps) message() string {
	return fmt.Sprintf("It's still %f Km away.", g.distance(g.location1, g.location2))
}

func main() {
	var curiosity = rover{}
	curiosity.gps.location1 = location{lat: -4.5895, long: 137.4417, name: "Bradbury Landing"}
	curiosity.gps.location2 = location{"Elysium Planitia", 4.5, 135.9}
	curiosity.gps.world = world{radius: 3389.5}
	fmt.Println(curiosity.message())
	fmt.Println(curiosity.location1.description())
}
