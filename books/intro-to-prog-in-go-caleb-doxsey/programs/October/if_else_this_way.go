// if_else_this_way.go
package main

import "fmt"

func main() {
    true_value := true
    false_value := false
    if fmt.Println(true_value); false_value {
        fmt.Println("This will not print")
    }

    if fmt.Println(false_value); true_value {
        fmt.Println("This prints")
    }

    if ;true_value{
        fmt.Println("This prints again")
    }

    if ;false_value {
        fmt.Println("This will not print")
    }
}