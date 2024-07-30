package elon

import (
	"strconv"
)

// Drive simulates driving the car by increasing the distance traveled and
// decreasing the battery level.
//
// No parameters. No return value.
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// DisplayDistance returns a string that represents the distance driven by the
// car.
//
// No parameters. Returns a string representing the distance driven by the
// car.
func (c *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(c.distance) + " meters"
}

// DisplayBattery returns a formatted string representing the battery level of
// the Car. It does not take any parameters. It returns a string.
func (c *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(c.battery) + "%"
}

// CanFinish determines if the car can finish a given distance.
//
// The parameter distance is the distance the car needs to travel.
// The return type is a boolean indicating if the car can finish the distance.
func (c *Car) CanFinish(distance int) bool {
	batteryUsage := distance * c.batteryDrain / c.speed
	return batteryUsage <= c.battery
}
