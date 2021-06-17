package variables

import (
	"fmt"
	"testing"
)

// TestConvertFahrenheitToCelcius calls temperature_conversion.ConvertFahrenheitToCelcius with
// a number: Fahrenheit temperature, checking for a valid number: Celcius return value
func TestConvertFahrenheitToCelcius(t *testing.T) {
	temps_array := []struct {
		fahrenheit float64
		celcius    float64
	}{
		{fahrenheit: 100, celcius: 37.77778},
		{fahrenheit: 0, celcius: -17.77778},
		{fahrenheit: -100, celcius: -73.33333},
		{fahrenheit: 32.00000, celcius: 0},
	}

	for _, tc := range temps_array {
		t.Run(fmt.Sprintf("F%f to C%f", tc.fahrenheit, tc.celcius), func(t *testing.T) {
			celcius, err := ConvertFahrenheitToCelcius(tc.fahrenheit)

			if err != nil {
				t.Fatalf(fmt.Sprintf("Error in test %v", err))
			}

			if celcius != tc.celcius {
				t.Errorf("Given Fahrenheit = %f, returned celcius = %f, %v, want match for %f, nil", tc.fahrenheit, celcius, err, tc.celcius)
			}
		})
	}

}
