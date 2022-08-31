// functions10.go
package main

import "fmt"

func firstFunction() {
    fmt.Println("firstFunction executes")
}

func secondFunction() int{
    defer firstFunction()
    fmt.Println("secondFunction starts")
    panic("Intentional panic in secondFunction")
    return 10
}

func main() {
    fmt.Printf("Output returned from second function: %d\n", secondFunction())
    fmt.Println("main function executes")
}