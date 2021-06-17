package variables

import (
    "fmt"
)

func total(){
    var (
        input_temperature float64
        // output_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    // TODO: Get result from output function
    // output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    var celcius float64
    
}