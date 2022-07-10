package elon

import "fmt"

// Drive update the number of meters driven and the battery.
func (c *Car) Drive() {
	if c.battery > c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance return the distance.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

// DisplayBattery return the battery percentage.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%s", c.battery, "%")
}

// CanFinish returns true if the car can finish the race; otherwise, return false.
func (c Car) CanFinish(trackDistance int) bool {
	return (c.speed/c.batteryDrain)*c.battery >= trackDistance
}
