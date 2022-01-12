// mapping4.go

package main

import "fmt"

func main() {
	elements := map[string]map[string]string{
		"H": {
			"name":          "Hydrogen",
			"atomic_number": "1",
		},
		"He": {
			"name":          "Helium",
			"atomic_number": "2",
		},
	}
	fmt.Println(elements)
}
