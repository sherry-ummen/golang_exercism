package space

import (
	"math"
)

type Planet string

var secondsInOneEarthYear float64 = 31557600.0

var orbitalPeriod = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age some comment
func Age(seconds float64, planet Planet) float64 {
	secondsInPlanet := secondsInOneEarthYear * orbitalPeriod[planet]
	return math.Round((seconds/secondsInPlanet)*100) / 100
}
