// mapping2.go
package main

import "fmt"

func main() {
	string_to_int := map[string]int{}
	string_to_float := map[string]float32{}
	string_to_string := map[string]string{}

	value_string_to_int, ok_string_to_int := string_to_int["notkey"]

	fmt.Println("Printing not found int value:", value_string_to_int, ", lookup success value:", ok_string_to_int)

	value_string_to_float, ok_string_to_float := string_to_float["notkey"]

	fmt.Println("Printing not found float value:", value_string_to_float, ", lookup success value:", ok_string_to_float)

	value_string_to_string, ok_string_to_string := string_to_string["notkey"]

	fmt.Println("Printing not found string value:", value_string_to_string, ", lookup success value:", ok_string_to_string)

	string_to_int["key"] = 10

	if value_string_to_int2, ok_string_to_int2 := string_to_int["key"]; ok_string_to_int2 {
		fmt.Printf("Lookup successful, %v, %v\n", value_string_to_int2, ok_string_to_int2)
	}

	if value_string_to_int3, ok_string_to_int3 := string_to_int["notkey"]; ok_string_to_int3 {
		fmt.Println("Lookup unsuccessful, this will never print")
		fmt.Println(value_string_to_int3, ok_string_to_int3)
	}
	fmt.Println("Here in lookup unsuccessful land")

}
