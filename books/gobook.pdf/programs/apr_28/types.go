// types.go
package main

import "fmt"

func main() {
	// Types
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.0+1.0)
	fmt.Println("32132 x 42452 = ", 32132*42452)
	fmt.Println("Length of 'Hello World' = ", len("Hello World"))
	fmt.Println("Hello World"[2])
	fmt.Println("Hello" + "World")
	fmt.Println(true && true, true && false, true || true, true || false, !true)

}
