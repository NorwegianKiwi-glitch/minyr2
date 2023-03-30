package yr

// Convert Celsius to Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

// Convert Fahrenheit to Celsius
func FarhenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}
