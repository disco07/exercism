package cars

import "math"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate) * float64(successRate)) / (60 * 100))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	if dec := int(math.Trunc(float64(carsCount / 10))); dec != 0 {
		return uint(dec*95000 + (carsCount-dec*10)*10000)
	}
	return uint(carsCount * 10000)
}
