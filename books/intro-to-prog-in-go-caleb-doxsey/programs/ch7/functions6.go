// functions6.go
package main

import "fmt"

func main(){
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven())
    fmt.Println(nextEven())
}

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint){
        ret = i
        i += 2
        return
    }
}
