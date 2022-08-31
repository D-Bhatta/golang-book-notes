package main

import "fmt"

func main() {
	var uname string = "johnny"
	validateUser(&uname, nil)
	fmt.Println("Finished validating user")
}

func recoverValidateUser() {
	if r := recover(); r != nil {
		fmt.Println("recovered from:", r)
	}
}

func validateUser(username *string, email *string) bool {
	defer recoverValidateUser()
	if username == nil {
		panic("runtime error: username passed for validation cannot be nil")
	}
	if email == nil {
		panic("runtime error: email passed for validation cannot be nil")
	}
	fmt.Printf("%s:%s", *username, *email)
	fmt.Println("Successfully validated user")
	return true
}
