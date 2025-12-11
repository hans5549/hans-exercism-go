// Package greeting provides functionalities related to greetings and weather forecasts.
package greeting

var (

	// CurrentCondition represents the current weather condition.
	CurrentCondition string

	// CurrentLocation represents the current location for the weather forecast.
	CurrentLocation string
)

// Forecast returns a string describing the current weather condition in the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
