package variables

import (
    "fmt"
    "testing"
)

func TestConvertFeetToMeters(t *testing.T){
    feet_array := []struct{
        feet float64
        meters float64
    } {
        {1, 0.3048},
        {99, 30.1752},
        {100000000, 30480000},
        {999999999, 304799999.6952},
        {0.9999999, 0.304799969520000046951},
        {1.121254778984587, 0.34175845663450216128},
        {0.00000000000000000000000000000001, 3.048000000000000277e-33},
    }
    
    for _, tc := range feet_array {
        t.Run(fmt.Sprintf("F%f to M%f", tc.feet, tc.meters), func(t *testing.T) {
            meters, err := ConvertFeetToMeters(tc.feet)

            if (err != nil) {
                t.Fatalf("Error in test %v", err)
            }

            if (meters != tc.meters){
                t.Errorf("Given feet=%f, returned meters=%.72f, %v, want match for %.72f,nil", tc.feet, meters, err, tc.meters)
            }
        })
    }
}