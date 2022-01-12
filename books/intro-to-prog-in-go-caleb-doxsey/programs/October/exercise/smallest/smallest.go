// smallest.go

package smallest

import (
	"errors"
	"fmt"
)

var ErrEmptyList = errors.New("list is empty: supply a list with at least one member")

func main(){
    fmt.Println("Printing smallest in the list:")
    list := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97, 9,17}
    fmt.Printf("48,96,86,68,57,82,63,70,37,34,83,27,19,97, 9,17")
    smallest, err := Smallest(list)
    if err == nil{
        fmt.Printf("Smallest in the list is %d", smallest)
    }
}

// returns the smallest value in an unsorted slice of ints
func Smallest(numlist []int) (int, error){
    if len(numlist) == 0 {
        return 0, ErrEmptyList
    }
    smallest := numlist[0]
    for _, num := range numlist{
        if num < smallest{
            smallest = num
        }
    }
    return smallest, nil
}