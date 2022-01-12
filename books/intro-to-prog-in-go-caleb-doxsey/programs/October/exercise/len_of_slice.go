// len_of_slice.go

package main

import "fmt"

func main(){
    fmt.Println(len(make([]int, 3, 9)))
    fmt.Println(make([]int, 3, 9))
}