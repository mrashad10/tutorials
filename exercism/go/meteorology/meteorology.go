package meteorology

import "fmt"

// TemperatureUnit represents the unit of temperature.
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String returns the string representation of the TemperatureUnit.
func (u TemperatureUnit) String() string {
	switch u {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return ""
	}
}

// Temperature represents a temperature value with a unit.
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String returns the string representation of the Temperature.
func (t Temperature) String() string {
	// Format the temperature as "<degree> <unit>"
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

// SpeedUnit represents the unit of speed.
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String returns the string representation of the SpeedUnit. It takes a
// SpeedUnit value as input and returns the corresponding string
// representation.
func (u SpeedUnit) String() string {
	// Use a switch statement to handle different SpeedUnit values
	switch u {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return ""
	}
}

// Speed represents a speed value with a unit.
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String returns the string representation of the Speed.
//
// It formats the magnitude and unit of the Speed as a string. The string is
// formatted as "<magnitude> <unit>", e.g. "10 mph".
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

// MeteorologyData represents meteorological data including location,
// temperature, wind direction, wind speed, and humidity.
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String returns the string representation of the MeteorologyData.
func (m MeteorologyData) String() string {
	// Format the MeteorologyData fields into a string using fmt.Sprintf.
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
