package main

import "fmt"

func main() {
	var t float64
	var ts string
	println("Please enter a temperature value:")
	fmt.Scanln(&t)

	println("Please enter a temperature scale to convert:")
	fmt.Scanln(&ts)

}

func convertTemperature(t float64, ts string) {

}

func celciusToFahrenheit(t float64) float64 {
	return t*9/5 + 32
}

func celciusToKelvin(t float64) float64 {
	return t + 273.15
}

func fahrenheitToCelcius(t float64) float64 {
	return (t - 32) * 5 / 9
}

// func fahrenheitToKelvin(t float64) float64 {
// 	return
// }

// func kelvinToCelcius(t float64) float64 {
// 	return
// }

// func kelvinTofahrenheit(t float64) float64 {
// 	return
// }

// Celsius to Fahrenheit: F=C×95+32F=C×59​+32
// Celsius to Kelvin: K=C+273.15K=C+273.15
// Fahrenheit to Celsius: C=(F−32)×59C=(F−32)×95​
// Fahrenheit to Kelvin: K=(F−32)×59+273.15K=(F−32)×95​+273.15
// Kelvin to Celsius: C=K−273.15C=K−273.15
// Kelvin to Fahrenheit: F=(K−273.15)×95+32F=(K−273.15)×59​+32
