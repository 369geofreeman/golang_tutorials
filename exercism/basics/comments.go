// Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition as a string.
var CurrentCondition string

// CurrentLocation represents the current location as a string.
var CurrentLocation string

// Forecast returns a string value representing the current weather condition in the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
