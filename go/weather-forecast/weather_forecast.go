// Package weather provides temperatures of country.
package weather

// CurrentCondition It has condition of weather.
var CurrentCondition string

// CurrentLocation It has location of city.
var CurrentLocation string

// Forecast returns a string of current location and
// current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
