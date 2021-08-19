// slice_1.go
package main

import "fmt"

func main() {
	var first_slice []float64
	fmt.Println("first_slice", first_slice)
	second_slice := make([]float64, 5, 10)
	fmt.Println("second_slice", second_slice)
	array_1 := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array_1", array_1)
	third_slice := array_1[0:5]
	fmt.Println("third_slice", third_slice)
	fourth_slice := []int{1, 2, 3, 4, 5}
	fmt.Println("fourth_slice", fourth_slice)
	fifth_slice := append(fourth_slice, 98, 99)
	fmt.Println("fifth_slice", fifth_slice)
	sixth_slice := make([]int, 2)
	copy(sixth_slice, fourth_slice)
	fmt.Println("sixth_slice", sixth_slice)
	seventh_slice := make([]int, 20)
	copy(seventh_slice, fourth_slice)
	fmt.Println("seventh_slice", seventh_slice)

}
