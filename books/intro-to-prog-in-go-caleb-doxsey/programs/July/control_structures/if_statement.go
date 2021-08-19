package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d is odd\n", i)
		} else {
			fmt.Println("This statement never prints")
		}
	}
}
