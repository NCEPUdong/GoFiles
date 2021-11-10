package main

import (
	"fmt"
	"math"
)

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere
type (
	coordinate struct {
		d, m, s float64
		h       rune
	}

	location struct {
		lat, long float64
	}

	landscape struct {
		roverOrLander string
		landingSite   string
		location
	}

	world struct {
		radius float64
	}

	distanceToSite struct {
		siteToSite string
		distance   float64
	}
)

// transfer rad to decimal
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// create a new location from coordinate
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// distance calculation using the Spherical Law of Cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radius
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// crete a list for distance between every two sites
func newDistanceList(landscapes []landscape, w world) []distanceToSite {
	distanceList := []distanceToSite{}
	for i := 0; i < len(landscapes); i++ {
		for j := i + 1; j < len(landscapes); j++ {
			siteToSite := landscapes[i].roverOrLander + "<=>" + landscapes[j].roverOrLander
			distance := w.distance(landscapes[i].location, landscapes[j].location)
			distanceList = append(distanceList, distanceToSite{siteToSite, distance})
		}
	}
	return distanceList
}

func maxDistance(d []distanceToSite) distanceToSite {
	var tmp = distanceToSite{siteToSite: d[0].siteToSite, distance: d[0].distance}
	for i := 0; i < len(d); i++ {
		if tmp.distance < d[i].distance {
			tmp.distance = d[i].distance
			tmp.siteToSite = d[i].siteToSite
		}
	}
	return tmp
}

func minDistance(d []distanceToSite) distanceToSite {
	var tmp = distanceToSite{siteToSite: d[0].siteToSite, distance: d[0].distance}
	for i := 0; i < len(d); i++ {
		if tmp.distance > d[i].distance {
			tmp.distance = d[i].distance
			tmp.siteToSite = d[i].siteToSite
		}
	}
	return tmp
}

func main() {
	// Bradbury Landing: 4°35'22.2" S, 137°26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	locations := []location{
		newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'}),
		newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'}),
		newLocation(lat, long),
		newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'}),
	}
	landscapes := []landscape{
		{roverOrLander: "Spirit", landingSite: "Columbia Memorial Station", location: locations[0]},
		{roverOrLander: "Opportunity", landingSite: "Challenger Memorial Station", location: locations[1]},
		{roverOrLander: "Curiosity", landingSite: "Bradbury Landing", location: locations[2]},
		{roverOrLander: "Insight", landingSite: "Elysium Planitia", location: locations[3]},
	}

	// for _, site := range landscapes {
	// 	fmt.Println(site)
	// }
	fmt.Println(landscapes)

	var mars = world{radius: 3389.5}
	distanceList := newDistanceList(landscapes, mars)
	fmt.Println(distanceList)

	maxDis := maxDistance(distanceList)
	minDis := minDistance(distanceList)
	fmt.Println("Maximun Distance is", maxDis)
	fmt.Println("Minimun Distance is", minDis)
}
