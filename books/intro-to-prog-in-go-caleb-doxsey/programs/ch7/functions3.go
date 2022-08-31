// functions3.go
package main

import "fmt"

func main() {
	NamedReturnFunction()
	first, second := MultipleReturnedValues()
	fmt.Println(first, second)
}

func NamedReturnFunction() (returned_value int) {
	returned_value = 1
	return
}

func MultipleReturnedValues() (int, int) {
	return 5, 6
}
