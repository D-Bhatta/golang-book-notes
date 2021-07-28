package ctrlstructs

import (
	"fmt"
	"testing"
)

func TestDivBy3Util(t *testing.T) {
	numbers_array := []struct {
		end     int
		numbers []int
	}{
		{0, []int{0}},
		{10, []int{0, 3, 6, 9}},
		{100, []int{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}},
	}

	for _, tc := range numbers_array {
		t.Run(fmt.Sprintf("E%d", tc.end), func(t *testing.T) {
			numbers, err := DivBy3Util(tc.end)

			if err != nil {
				t.Fatalf("Error in test %v", err)
			}

			if len((numbers)) != len(tc.numbers) {
				t.Errorf("Given end=%d, returned len(numbers)=%d, %v, want match for len(numbers)=%d,nil", tc.end, len(numbers), err, len(tc.numbers))
			} else {
				for i := 0; i < len(tc.numbers); i++ {
					if numbers[i] != tc.numbers[i] {
						t.Errorf("Given end=%d, returned numbers[%d]=%d, %v, want match for numbers[%d]=%d,nil", tc.end, i, numbers[i], err, i, tc.numbers[i])
					}
				}
			}

		})
	}
}
