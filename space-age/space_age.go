package space

// Planet one of the solar system planets
type Planet string

var ageConvertMap map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

var earthSecondsInAYear float64 = 31557600

// Age : given seconds and planet, returns the age on that planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (earthSecondsInAYear * ageConvertMap[planet])
}
