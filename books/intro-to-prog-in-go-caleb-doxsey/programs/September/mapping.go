// mapping.go
package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["key"] = 10
	fmt.Println(x)

	y := make(map[int]int)
	y[0] = 0
	y[1] = 1
	fmt.Println(y)
	fmt.Println(len(y))
	delete(y, 0)
	fmt.Println(y)
	fmt.Println(len(y))

	z := map[string]float32{}

	z["red"] = 0.01
	fmt.Println(z)
}
