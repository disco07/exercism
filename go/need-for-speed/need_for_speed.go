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
	car = Car{speed: car.speed, batteryDrain: car.batteryDrain, distance: car.distance, battery: car.battery}
	if car.batteryDrain < car.battery {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	car = NewCar(car.speed, car.batteryDrain)
	track = NewTrack(track.distance)

	if car.battery-car.batteryDrain < track.distance {
		return false
	}
	return true
}
