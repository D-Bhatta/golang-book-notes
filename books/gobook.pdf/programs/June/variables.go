package main

import "fmt"

func main() {

	// Variables

	var x string = "Hello World"
	var y string
	y = "New World"
	fmt.Println(x, y)
	x += y
	fmt.Println(x)

	var hello_string string = "Hello"
	var world_string string = "World"

	fmt.Println(hello_string == world_string)

}
