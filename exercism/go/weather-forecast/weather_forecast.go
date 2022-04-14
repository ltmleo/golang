// Package weather is a weather forecast package.
package weather 

// CurrentCondition stores the current weather condition.
var CurrentCondition string 
// CurrentLocation stores the current location.
var CurrentLocation string 

// Forecast function return the current weather for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
