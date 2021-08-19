package variables

import (
	"fmt"
	"log"
	"math"
)

func total() float64{
    var (
        input_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    output_temperature, err := ConvertFahrenheitToCelcius(input_temperature)

    if (err != nil) {
        log.Fatal(fmt.Sprintf("Error in function total in variables/temperature_conversion.go: %v", err))
    }

    return output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    celcius := ((fahrenheit_temp-32) * 5)/9
    celcius_6f := math.Round((celcius * 1000000))/1000000
    return celcius_6f, nil
}