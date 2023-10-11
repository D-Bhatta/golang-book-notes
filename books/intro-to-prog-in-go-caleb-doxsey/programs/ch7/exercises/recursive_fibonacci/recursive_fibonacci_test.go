// recursive_fibonacci_test.go

package recursivefibonacci

import (
	"fmt"
	"testing"
)

func TestRecursiveFibonacci(t *testing.T) {
	numbers := []struct {
		ID           uint
		InputNumber  uint64
		OutputNumber uint64
	}{
		{1, 0, 0},
		{2, 1, 1},
        {3, 2, 1},
        {4, 10, 55},
        {5, 15, 610},
        {6, 30, 832040},
	}

	for _, testCase := range numbers {
		t.Run(fmt.Sprintf("Test Case: %d:", testCase.ID), func(t *testing.T) {
			outputNumber := RecursiveFibonacci(testCase.InputNumber)
		
			if outputNumber != testCase.OutputNumber {
				t.Errorf("Returned output number: %d does not match expected output number: %d for input number %d.", outputNumber, testCase.OutputNumber, testCase.InputNumber)
			}
		})
	}
}
