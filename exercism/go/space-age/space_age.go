package space

type Planet string

// earthYearInSeconds is a constant that represents the number of seconds in
// one Earth year.
const earthYearInSeconds = 31557600.0

// conversionRates is a map that stores the conversion rates from Earth years
// to years on other planets. The keys are the names of the planets and the
// values are the conversion rates.
var conversionRates = map[Planet]float64{
	"Earth":   earthYearInSeconds,
	"Mercury": earthYearInSeconds * 0.2408467,
	"Venus":   earthYearInSeconds * 0.61519726,
	"Mars":    earthYearInSeconds * 1.8808158,
	"Jupiter": earthYearInSeconds * 11.862615,
	"Saturn":  earthYearInSeconds * 29.447498,
	"Uranus":  earthYearInSeconds * 84.016846,
	"Neptune": earthYearInSeconds * 164.79132,
}

// Age calculates the age in years based on the number of seconds and the
// given planet.
//
// The function takes two parameters: - seconds: a float64 representing the
// number of seconds - planet: a Planet type representing the planet on which
// the age is calculated
//
// The function returns a float64 representing the age in years. If the
// conversion rate for the given planet is not found, the function returns
// -1.0.
func Age(seconds float64, planet Planet) float64 {
	rate, ok := conversionRates[planet]
	if !ok {
		return -1.0
	}

	return seconds / rate
}
