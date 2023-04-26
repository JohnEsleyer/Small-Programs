package main

import "fmt"

func main() {
	var temperature float64
	var unit string

	fmt.Print("Enter the temperature: ")
	fmt.Scanln(&temperature)

	fmt.Print("Enter the unit of temperature (C or F): ")
	fmt.Scanln(&unit)

	if unit == "C" {
		// convert Celsius to Fahrenheit
		fahrenheit := (temperature * 9 / 5) + 32
		fmt.Printf("%.2f°C is %.2f°F\n", temperature, fahrenheit)
	} else if unit == "F" {
		// convert Fahrenheit to Celsius
		celsius := (temperature - 32) * 5 / 9
		fmt.Printf("%.2f°F is %.2f°C\n", temperature, celsius)
	} else {
		fmt.Println("Invalid unit of temperature. Please enter C or F.")
	}
}
