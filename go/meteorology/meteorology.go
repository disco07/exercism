package meteorology

import "fmt"

type TemperatureUnit int

type Stringer interface {
	String() string
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String method to the TemperatureUnit type.
func (t TemperatureUnit) String() string {
	if t == 0 {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String method to the Temperature type.
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	if s == 0 {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %v, Wind %s at %v, %v%% Humidity", m.location, m.temperature.String(), m.windDirection, m.windSpeed.String(), m.humidity)
}
