package collections

import (
	"encoding/json"
	"fmt"
	"math"
)

// TestStruct test work with srtucture
func TestStruct() {
	type location struct {
		name string
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{name: "Columbia Station", Lat: -14.5684, Long: 175.472636},
		{name: "Challenger Station", Lat: -1.9462, Long: 354.4734},
	}
	// for _, loc := range locations {
	bytes, err := json.MarshalIndent(locations, "", " ")
	if err != nil {
		fmt.Printf("error convert to json %v\n", locations)
		// continue
	}
	fmt.Println(string(bytes))
	// }
}

type world struct {
	radius float64
}

type coordinateDMS struct {
	d, m, s float64
	h       rune
}

type coordinateDD struct {
	deg float64
}

type location struct {
	lat  float64
	long float64
	name string
}

func (c coordinateDMS) toDecimal() float64 {
	sign := 1.0
	switch c.h {
	case 'W', 'w', 'S', 's':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(name string, lat, long float64) location {
	return location{
		name: name,
		lat:  lat,
		long: long,
	}
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// TestStructModules test modules with srtucture
func TestStructModules() {

	lat := coordinateDMS{4, 35, 22.2, 'S'}
	long := coordinateDMS{137, 26, 30.1, 'E'}
	curiosity := newLocation("Curiosity", lat.toDecimal(), long.toDecimal())

	lat = coordinateDMS{14, 34, 36.2, 'S'}
	long = coordinateDMS{175, 28, 21.5, 'E'}
	spirit := newLocation("Spirit", lat.toDecimal(), long.toDecimal())

	lat = coordinateDMS{1, 56, 46.3, 'S'}
	long = coordinateDMS{354, 28, 24.2, 'E'}
	opportunity := newLocation("Opportunity", lat.toDecimal(), long.toDecimal())

	lat = coordinateDMS{4, 30, 0.0, 'N'}
	long = coordinateDMS{135, 54, 0.0, 'E'}
	inSight := newLocation("InSight", lat.toDecimal(), long.toDecimal())

	landings := []location{
		spirit,
		curiosity,
		opportunity,
		inSight,
	}

	fmt.Println(landings)

	mars := world{radius: 3389.5}

	distances := []float64{
		mars.distance(curiosity, spirit),
		mars.distance(curiosity, opportunity),
		mars.distance(curiosity, inSight),
		mars.distance(spirit, opportunity),
		mars.distance(spirit, inSight),
		mars.distance(opportunity, inSight),
	}

	minDist := 0.0
	maxDist := 0.0
	for i, val := range distances {
		if i == 0 {
			minDist = val
			maxDist = val
			continue
		}
		if minDist > val {
			minDist = val
		}
		if maxDist < val {
			maxDist = val
		}
	}
	fmt.Println(distances)
	fmt.Printf("maximum dist - %v\nminimum dist - %v", maxDist, minDist)
	//
	// spirit = newLocation(-14.5684, 175.472636)
	// oport := newLocation(-1.9462, 354.4734)
	// dist := mars.distance(spirit, oport)
	// fmt.Printf("%.2f km\n", dist)

}
