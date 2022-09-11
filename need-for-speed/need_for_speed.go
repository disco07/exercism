package speed

type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain <= car.battery {
		car.distance += car.speed
		car.battery -= car.batteryDrain
		return car
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var maxDist float64
	maxDist = (float64(car.speed) / float64(car.batteryDrain)) * float64(car.battery)
	if maxDist+float64(car.distance) > float64(track.distance) {
		return true
	} else {
		return false
	}
}
