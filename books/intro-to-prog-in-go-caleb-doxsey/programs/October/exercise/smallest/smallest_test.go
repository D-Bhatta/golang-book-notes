// smallest_test.go

package smallest

import (
	"fmt"
	"testing"
)

func TestSmallest(t *testing.T) {
	num_array := []struct {
		nums         []int
		smallest     int
		err_expected error
	}{
		{[]int{1}, 1, nil},
		{[]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2}, -1, nil},
		{[]int{90, 30, 12, 12, 1200}, 12, nil},
		{[]int{4, 1, 1, 0, 2, 8}, 0, nil},
		{[]int{}, 0, ErrEmptyList}, // 0 is a placeholder, can be replaced with whatever, it is never used.
	}

	for _, tc := range num_array {
		t.Run(fmt.Sprintf("N%d", tc.smallest), func(t *testing.T) {
			smallest, err := Smallest(tc.nums)
			if err != tc.err_expected {
				t.Fatalf("Error in test %v. Expected error is %v.", err, tc.err_expected)
			}

			if smallest != tc.smallest {
                t.Errorf("Returned smallest = %d, %v. Want match for %d, %v.", smallest, err, tc.smallest, tc.err_expected)
			}
		})
	}
}

var smallest_result int

func BenchmarkSmallest(b *testing.B){
    var result int
    for n := 0; n <b.N; n++{
        result, _ = Smallest([]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2})
    }
    smallest_result = result
}
