// fizz_buzz.go
package ctrlstructs

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidInput = errors.New("input is less than 0")

func PrintFizzBuzz() {
	// TODO: Call FizzBuzz and print the data
	fmt.Print("")
}

// Returns a string slice from 1 to n of FizzBuzz
// n must be 0 or greater
func FizzBuzzUtil(n int) ([]string, error) {
	if n < 0 {
		return []string{}, ErrInvalidInput
	}

	fizz_buzz := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizz_buzz = append(fizz_buzz, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fizz_buzz = append(fizz_buzz, "Fizz")
			continue
		}
		if i%5 == 0 {
			fizz_buzz = append(fizz_buzz, "Buzz")
			continue
		}
		fizz_buzz = append(fizz_buzz, strconv.Itoa(i))
	}

	return fizz_buzz, nil
}
