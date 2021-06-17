package main

import (
	"fmt"
    "log"
)


func new() {
    fmt.Print("Enter an number: ")
    var number_input float64
    _, err := fmt.Scanf("%f", &number_input)
    if err != nil {
        log.Fatal(err)
    }

    mult_result := number_input * 2

    fmt.Println("Result is: ", mult_result)
}