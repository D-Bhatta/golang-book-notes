// functions4.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(add(1,2,3))
    nums := []int{1,2,3,4,5,6,7,8,9,10,0}
    fmt.Println(add(nums...))
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
    fmt.Println(reflect.TypeOf(args))
	return total
}