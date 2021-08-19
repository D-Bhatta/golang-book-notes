// array_1.go
package main

import "fmt"

func main() {
	var numbers [5]float64

	numbers[0] = 98

	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println(total / float64(len(numbers)))

    var other_total float64= 0
    for _,value := range numbers {
        other_total += value
    }
    fmt.Println(other_total/float64(len(numbers)))

    second_array := [5]float64{ 98, 0, 0, 10, 11}
    fmt.Println(second_array)

    third_array := [3]string{
        "hello",
        "I",
        "am",
    }

    fmt.Println(third_array)
}
