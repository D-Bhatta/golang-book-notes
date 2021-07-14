// length_conversion.go
package variables

import (
	"fmt"
	"log"
)

func FeetToMeters() int {
	var input_number float64
	fmt.Print("Enter the value of feet: ")
	fmt.Scanf("%f", &input_number)

	if input_number < 0 {
		log.Fatalf("%f is not a valid input: The value of feet cannot be less than 0, please input a non-negative number", input_number)
		return 1
	}

	meters, err := ConvertFeetToMeters(input_number)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error in function FeetToMeters in variables/length_conversion.go: %v", err))
	}

	fmt.Printf("%f feet is %f meters", input_number, meters)

	return 0
}

func ConvertFeetToMeters(feet float64) (float64, error) {
	const constant_meters float64 = 0.30480
	var meters float64 = feet * constant_meters

	return meters, nil
}
