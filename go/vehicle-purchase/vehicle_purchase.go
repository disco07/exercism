package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicle := strings.Split(option1, " ")
	vehicle2 := strings.Split(option2, " ")
	for i := 0; i < len(vehicle); i++ {
		if vehicle[i] > vehicle2[i] {
			return fmt.Sprintf("%s is clearly the better choice.", option2)
		} else if vehicle[i] < vehicle2[i] {
			return fmt.Sprintf("%s is clearly the better choice.", option1)
		}
	}
	return ""
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3.0:
		return (originalPrice * 80) / 100
	case age >= 10.0:
		return (originalPrice * 50) / 100
	case age >= 3 && age < 10:
		return (originalPrice * 70) / 100
	default:
		return originalPrice
	}
}
