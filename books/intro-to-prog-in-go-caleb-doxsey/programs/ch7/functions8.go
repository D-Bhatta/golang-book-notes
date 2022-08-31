// functions8.go
package main

import "fmt"

func firstFunction() {
	fmt.Println("First function executes")
}

func secondFunction() {
	fmt.Println("Second function executes")
}

func main() {
	defer firstFunction()
	secondFunction()
}
