// length_conversion.go
package variables

import (
	"fmt"
)

func FeetToMeters() float64 {
	var input_number float64
	fmt.Print("Enter the value of feet: ")
	fmt.Scanf("%f", &input_number)
    // TODO: Check for negative values

	// TODO: Get converted meters from feet value using function
	return 1
}

func ConvertFeetToMeters(feet float64) (float64, error) {
    const constant_meters float64 = 0.30480
    var meters float64 = feet * constant_meters

    return meters,nil
}
