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
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

// CanFinish returns true if the car can finish the race; otherwise, return false.
func (c Car) CanFinish(trackDistance int) bool {
	fmt.Println((c.speed / c.batteryDrain) * c.battery)
	return (float64(c.speed)/float64(c.batteryDrain))*float64(c.battery) >= float64(trackDistance)
}
