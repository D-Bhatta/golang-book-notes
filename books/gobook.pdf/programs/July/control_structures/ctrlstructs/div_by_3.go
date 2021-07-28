package ctrlstructs

import (
	"errors"
	"fmt"
	"log"
)

func DivBy3Printer() {
	numbers, err := DivBy3Util(100)
	if err != nil {
        log.Fatal(fmt.Sprintf("Error in function DivBy3Printer in ctrlstructs/div_by_3.go: %v", err))
	}
	fmt.Print(numbers)

}

/* DivBy3Util returns an array of numbers that are divisible by 3 from 0 till end.
If the end variable is less than 0, return an error
*/
func DivBy3Util(end int) ([]int, error) {
	if end < 0 {
		return []int{}, errors.New("passed parameter end cannot be a negative value")
	}

	var numbers []int

	for i := 0; i <= end; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}

	}

	return numbers, nil
}
