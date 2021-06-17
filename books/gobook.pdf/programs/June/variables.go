// variables.go
package main

import "fmt"

var global_variable string = "This is a global variable"

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

	new_string := "This is a string whose type will be inferred"
	fmt.Println(new_string)

    fmt.Println(global_variable)

    var (
        intialized_a = 5
        intialized_b = 10
        intialized_c = 15
    )

    const (
        intialized_x = 50
        intialized_y = 100
        intialized_z = 150
    )

    fmt.Println(intialized_a, intialized_b, intialized_c)
    fmt.Println(intialized_x, intialized_y, intialized_z)
}
