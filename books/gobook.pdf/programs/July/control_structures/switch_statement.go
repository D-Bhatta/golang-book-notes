package main

import "fmt"

func main() {
	var input int

	fmt.Print("Enter a number between 1 to 4: ")
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		{
			fmt.Println("Choses case 1")
		}
	case 2:
		{
			fmt.Println("Choses case 2")
		}
	case 3:
		{
			fmt.Println("Choses case 3")
		}
	case 4:
		{
			fmt.Println("Choses case 4")
		}
	default:
		fmt.Println("Chosen number is out of range")
	}
}
