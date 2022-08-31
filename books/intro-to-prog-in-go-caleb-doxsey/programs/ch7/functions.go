// functions.go
package main

import "fmt"

func main() {
	numbers := [5]float64{98, 93, 77, 82, 83}

	fmt.Println(average(numbers[:]))
}

func average(numbers []float64) float64 {
	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total / float64(len(numbers))
}
