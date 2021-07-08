# Go Book

Book by Caleb Doxsey.

<!-- markdownlint-disable MD010 -->

## Sections

- [Go Book](#go-book)
  - [Sections](#sections)
  - [Notes](#notes)
  - [Go basics](#go-basics)
    - [Running go files](#running-go-files)
    - [Go programs](#go-programs)
    - [Types in go](#types-in-go)
      - [Numbers](#numbers)
      - [Variables](#variables)
    - [Exercise Programs](#exercise-programs)
      - [Temparature conversion from Fahrenheit to C](#temparature-conversion-from-fahrenheit-to-c)
      - [Convert from feet to meters](#convert-from-feet-to-meters)
  - [Additional notes](#additional-notes)
    - [Language specific](#language-specific)

## Notes

## Go basics

### Running go files

The `go run` command takes the subsequent files (separated by spaces), compiles them into an executable saved in a temporary directory and then runs the program.

### Go programs

Every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code.
There are two types of Go programs: executables and libraries. Executable applications are the kinds of programs that we can run directly from the terminal (in Windows they end with .exe).
Libraries are collections of code that we package together so that we can use them in other programs.

```powershell
go run main.go
```

The `import` keyword is how we include code from other packages to use with our program. The `fmt` package
(shorthand for format) implements formatting for input and output.

Go supports two different styles of comments: `// comments` in which all the text between the // and the end of the line is part of the comment and `/* */ comments` where everything between the *s is part of the comment.

Functions are the building blocks of a Go program. They have inputs, outputs and a series of steps called
statements which are executed in order. All functions start with the keyword `func` followed by the name of
the function (`main` for example), a list of zero or more parameters surrounded by parentheses, an optional
return type and a body which is surrounded by curly braces. The name `main` is special because it's the function that gets called when you execute the program.

To get help from `godoc` use

```powershell
godoc fmt Println
```

### Types in go

Go is statically typed. This removes entire set of problems.

#### Numbers

Go's integer types are: `uint8, uint16, uint32, uint64, int8, int16, int32 and int64`. **uint** means *un-signed integer* which means non-negative. **int** means *signed integer*. Un-signed integers only contain positive numbers (or zero).
Alias types: `byte` which is the same as `uint8` and `rune` which is the same as `int32`.

Machine dependent integer types are `uint`, `int`, and `uintptr` whose size depends on the machine architecture being used.

Floating point numbers are inexact. They are `float32` and `float64`. For complex numbers we have `complex64` and `complex128`.

Strings are represented using `""` or ` `` `

Double quoted strings allow special escape characters, and they cannot contain new lines.

Length is found using len(), random access is allowed via [], and concatenation via +.

Logical operators are

- && and
- || or
- ! not

```go
// types.go
package main

import "fmt"

func main() {
	// Types
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.0+1.0)
	fmt.Println("32132 x 42452 = ", 32132*42452)
	fmt.Println("Length of 'Hello World' = ", len("Hello World"))
	fmt.Println("Hello World"[2])
	fmt.Println("Hello" + "World")
	fmt.Println(true && true, true && false, true || true, true || false, !true)

}

```

```powershell
gobook.pdf/programs/apr_28 on  gobook [!⇡] via 🐹 v1.16.4 on ☁️  (us-east-1) 
❯ go run .\types.go
1 + 1 =  2
1 + 1 =  2
32132 x 42452 =  1364067664  
Length of 'Hello World' =  11
108
HelloWorld
true false true true false
```

#### Variables

- We declare variables in the form of `var <name> <type> = <value>`.
- They can be created first and initialized later.
- Variables created with `var` can be reassigned.
- There exists the `+=` operator which expands to `x = x + y` for `x += y`.
- Equality operator is `==`.
- Variable with a starting value and no type specified can be created with `:=`.
- This should be used whenever possible.
- Variables can be named outside of all scoped structures, and be globally accessible.
- To define multiple variables, using `var` and `const`, enclose them ().
- Use `fmt.Print` to print something in the same line.
- `fmt.Scanf` is used to read the user input into a memory location.

```go
// variables.go
package main

import "fmt"

var global_variable string = "This is a global variable"

func main() {

	// Variables

	var x string = "Hello World"
	var y string
	y = "New World"
	fmt.Println(x, y)
	x += y
	fmt.Println(x)

	var hello_string string = "Hello"
	var world_string string = "World"

	fmt.Println(hello_string == world_string)

	new_string := "This is a string whose type will be inferred"
	fmt.Println(new_string)

    fmt.Println(global_variable)

    var (
        intialized_a = 5
        intialized_b = 10
        intialized_c = 15
    )

    const (
        intialized_x = 50
        intialized_y = 100
        intialized_z = 150
    )

    fmt.Println(intialized_a, intialized_b, intialized_c)
    fmt.Println(intialized_x, intialized_y, intialized_z)
}

```

- Example program [mult_by_2.go](./programs/June/mult_by_2.go)

```go
package main

import (
	"fmt"
    "log"
)


func new() {
    fmt.Print("Enter an number: ")
    var number_input float64
    _, err := fmt.Scanf("%f", &number_input)
    if err != nil {
        log.Fatal(err)
    }

    mult_result := number_input * 2

    fmt.Println("Result is: ", mult_result)
}
```

### Exercise Programs

First we create a module `go mod init debabrata.xyz/variables`.

#### Temparature conversion from Fahrenheit to C

```txt
Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
```

- Create a file `temperature_conversion.go`
- Add skeletal structure

```go
package variables

import (
    "fmt"
)

func total(){
    var (
        input_temperature float64
        // output_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    // TODO: Get result from output function
    // output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    return 1, nil
}
```

- Create test file `temperature_conversion_test.go`

```go
package variables

import (
	"fmt"
	"testing"
)

// TestConvertFahrenheitToCelcius calls temperature_conversion.ConvertFahrenheitToCelcius with
// a number: Fahrenheit temperature, checking for a valid number: Celcius return value
func TestConvertFahrenheitToCelcius(t *testing.T) {
	temps_array := []struct {
		fahrenheit float64
		celcius    float64
	}{
		{fahrenheit: 100.000000, celcius: 37.777778},
		{fahrenheit: 0.000000, celcius: -17.777778},
		{fahrenheit: -100.000000, celcius: -73.333333},
		{fahrenheit: 32.000000, celcius: 0},
	}

	for _, tc := range temps_array {
		t.Run(fmt.Sprintf("F%f to C%f", tc.fahrenheit, tc.celcius), func(t *testing.T) {
			celcius, err := ConvertFahrenheitToCelcius(tc.fahrenheit)

			if err != nil {
				t.Fatalf("Error in test %v", err)
			}

			if celcius != tc.celcius {
				t.Errorf("Given Fahrenheit = %f, returned celcius = %f, %v, want match for %f, nil", tc.fahrenheit, celcius, err, tc.celcius)
			}
		})
	}

}

```

- The function `TestConvertFahrenheitToCelcius` is the test function which takes the argument `t` of type `*testing.T`
- We create an array `temps_array` of test cases with the `struct` of the following types

```go
struct {
    fahrenheit float64
    celcius    float64
}
```

- They can be accessed using the `.` operator.
- We initialize the array with test cases.
- Make sure the last line in the block has a `,` as well.
- Then we iterate over the test cases.
- We use subtests to run the tests concurrently using `t.Run` which takes a name `string` and a `func`.
- We add a name to each test with `fmt.Sprintf("F%f to C%f", tc.fahrenheit, tc.celcius)`.
- `fmt.Sprintf` returns a string formatted.
- We get the variable from the `ConvertFahrenheitToCelcius` as `celcius`.
- We check for error and return a fatal error with `t.Fatalf` which accepts a string.
- We match the result against the wanted test case result.
- If it isn't a match we use the `t.Errorf` to show failure.
- We run the tests with `go test -v`
- Since we are using TDD, we make sure that the tests fail first.

```powershell
programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯ go test -v
=== RUN   TestConvertFahrenheitToCelcius
=== RUN   TestConvertFahrenheitToCelcius/F100.000000_to_C37.777780
    temperature_conversion_test.go:30: Given Fahrenhiet = 100.000000, returned celcius = 1.000000, <nil>, want match for 37.777780, nil
=== RUN   TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777780
    temperature_conversion_test.go:30: Given Fahrenhiet = 0.000000, returned celcius = 1.000000, <nil>, want match for -17.777780, nil
=== RUN   TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333330
    temperature_conversion_test.go:30: Given Fahrenhiet = -100.000000, returned celcius = 1.000000, <nil>, want match for -73.333330, nil
=== RUN   TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000
    temperature_conversion_test.go:30: Given Fahrenhiet = 32.000000, returned celcius = 1.000000, <nil>, want match for 0.000000, nil
--- FAIL: TestConvertFahrenheitToCelcius (0.32s)
    --- FAIL: TestConvertFahrenheitToCelcius/F100.000000_to_C37.777780 (0.06s)
    --- FAIL: TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777780 (0.01s)
    --- FAIL: TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333330 (0.02s)
    --- FAIL: TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000 (0.03s)
FAIL
exit status 1
FAIL    debabrata.xyz/variables 0.604s

programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯
```

- We write the logic to convert fahrenheit to celcius values in the program

```go
package variables

import (
    "fmt"
)

func total(){
    var (
        input_temperature float64
        // output_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    // TODO: Get result from output function
    // output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    celcius := ((fahrenheit_temp-32) * 5)/9
    return celcius, nil
}
```

- We run tests with go test

```powershell
programs/June/variables on  gobook [!⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) 
❯ go test -v
=== RUN   TestConvertFahrenheitToCelcius
=== RUN   TestConvertFahrenheitToCelcius/F100.000000_to_C37.777778
    temperature_conversion_test.go:30: Given Fahrenheit = 100.000000, returned celcius = 37.777778, <nil>, want match for 37.777778, nil
=== RUN   TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777778
    temperature_conversion_test.go:30: Given Fahrenheit = 0.000000, returned celcius = -17.777778, <nil>, want match for -17.777778, nil
=== RUN   TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333333
    temperature_conversion_test.go:30: Given Fahrenheit = -100.000000, returned celcius = -73.333333, <nil>, want match for -73.333333, nil
=== RUN   TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000
--- FAIL: TestConvertFahrenheitToCelcius (0.07s)
    --- FAIL: TestConvertFahrenheitToCelcius/F100.000000_to_C37.777778 (0.00s)
    --- FAIL: TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777778 (0.00s)
    --- FAIL: TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333333 (0.07s)
    --- PASS: TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/variables 0.453s

programs/June/variables on  gobook [!⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 26s
❯
```

- We notice that one test passes and others fail. This is beacause of floating point numbers.
- We only need precision of upto 6 decimal digits.

```go
package variables

import (
	"fmt"
	"math"
)

func total(){
    var (
        input_temperature float64
        // output_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    // TODO: Get result from output function
    // output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    celcius := ((fahrenheit_temp-32) * 5)/9
    celcius_6f := math.Round((celcius * 1000000))/1000000
    return celcius_6f, nil
}
```

- We run tests with `go test -v`

```powershell
programs/June/variables on  gobook [!⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯ go test -v
=== RUN   TestConvertFahrenheitToCelcius
=== RUN   TestConvertFahrenheitToCelcius/F100.000000_to_C37.777778
=== RUN   TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777778
=== RUN   TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333333
=== RUN   TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000
--- PASS: TestConvertFahrenheitToCelcius (0.01s)
    --- PASS: TestConvertFahrenheitToCelcius/F100.000000_to_C37.777778 (0.00s)
    --- PASS: TestConvertFahrenheitToCelcius/F0.000000_to_C-17.777778 (0.00s)
    --- PASS: TestConvertFahrenheitToCelcius/F-100.000000_to_C-73.333333 (0.00s)
    --- PASS: TestConvertFahrenheitToCelcius/F32.000000_to_C0.000000 (0.00s)
PASS
ok      debabrata.xyz/variables 0.095s

programs/June/variables on  gobook [!⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯
```

- We polish off the program

```go
package variables

import (
	"fmt"
	"log"
	"math"
)

func total() float64{
    var (
        input_temperature float64
    )
    fmt.Print("Enter the input temperature in Fahrenheit: ")
    fmt.Scanf("%f", &input_temperature)

    output_temperature, err := ConvertFahrenheitToCelcius(input_temperature)

    if (err != nil) {
        log.Fatal(fmt.Sprintf("Error in function total in variables/temperature_conversion.go: %v", err))
    }

    return output_temperature
}

func ConvertFahrenheitToCelcius(fahrenheit_temp float64) (float64, error){
    celcius := ((fahrenheit_temp-32) * 5)/9
    celcius_6f := math.Round((celcius * 1000000))/1000000
    return celcius_6f, nil
}
```

#### Convert from feet to meters

- In this program we will convert feet to meters.
- We first create a skeleton file

```go
// length_conversion.go
package variables

import (
	"fmt"
)

func FeetToMeters() float64 {
	var input_number float64
	fmt.Print("Enter the value of feet: ")
	fmt.Scanf("%f", &input_number)

	// TODO: Get converted meters from feet value using function
	return 1
}

func ConvertFeetToMeters(feet float64) (float64, error) {
    return 1, nil
}

```

- We create a test file with tests

```go
package variables

import (
    "fmt"
    "testing"
)

func TestConvertFeetToMeters(t *testing.T){
    feet_array := []struct{
        feet float64
        meters float64
    } {
        {1, 0.3048},
        {99, 30.1752},
        {100000000, 30480000},
        {999999999, 304799999.695},
        {0.9999999, 0.30479996952},
        {1.121254778984587, 0.34175845663450216128},
        {0.00000000000000000000000000000001, 3.048000000000000277e-33},
    }
    
    for _, tc := range feet_array {
        t.Run(fmt.Sprintf("F%f to M%f", tc.feet, tc.meters), func(t *testing.T) {
            meters, err := ConvertFeetToMeters(tc.feet)

            if (err != nil) {
                t.Fatalf("Error in test %v", err)
            }

            if (meters != tc.meters){
                t.Errorf("Given feet=%f, returned meters=%f, %v, want match for %f,nil", tc.feet, meters, err, tc.meters)
            }
        })
    }
}
```

- We run the test with `go test -v -run TestConvertFeetToMeters`

```powershell
programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1)
❯ go test -v -run TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800
    length_conversion_test.go:31: Given feet=1.000000, returned meters=1.000000, <nil>, want match for 0.304800,nil
=== RUN   TestConvertFeetToMeters/F99.000000_to_M30.175200
    length_conversion_test.go:31: Given feet=99.000000, returned meters=1.000000, <nil>, want match for 30.175200,nil
=== RUN   TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000
    length_conversion_test.go:31: Given feet=100000000.000000, returned meters=1.000000, <nil>, want match for 30480000.000000,nil
=== RUN   TestConvertFeetToMeters/F999999999.000000_to_M304799999.695000
    length_conversion_test.go:31: Given feet=999999999.000000, returned meters=1.000000, <nil>, want match for 304799999.695000,nil
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800#01
    length_conversion_test.go:31: Given feet=1.000000, returned meters=1.000000, <nil>, want match for 0.304800,nil
=== RUN   TestConvertFeetToMeters/F1.121255_to_M0.341758
    length_conversion_test.go:31: Given feet=1.121255, returned meters=1.000000, <nil>, want match for 0.341758,nil
=== RUN   TestConvertFeetToMeters/F0.000000_to_M0.000000
    length_conversion_test.go:31: Given feet=0.000000, returned meters=1.000000, <nil>, want match for 0.000000,nil
--- FAIL: TestConvertFeetToMeters (0.06s)
    --- FAIL: TestConvertFeetToMeters/F1.000000_to_M0.304800 (0.00s)
    --- FAIL: TestConvertFeetToMeters/F99.000000_to_M30.175200 (0.00s)
    --- FAIL: TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000 (0.00s)
    --- FAIL: TestConvertFeetToMeters/F999999999.000000_to_M304799999.695000 (0.01s)
    --- FAIL: TestConvertFeetToMeters/F1.000000_to_M0.304800#01 (0.00s)
    --- FAIL: TestConvertFeetToMeters/F1.121255_to_M0.341758 (0.02s)
    --- FAIL: TestConvertFeetToMeters/F0.000000_to_M0.000000 (0.01s)
FAIL
exit status 1
FAIL    debabrata.xyz/variables 0.155s

programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1)
❯
```

- We add code to return a conversion to meters

```go
func ConvertFeetToMeters(feet float64) (float64, error) {
    meters := feet * 0.3048

    return meters,nil
}
```

- We run the test with `go test -v -run TestConvertFeetToMeters`

```powershell
programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯ go test -v -run TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800
=== RUN   TestConvertFeetToMeters/F99.000000_to_M30.175200
=== RUN   TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000
=== RUN   TestConvertFeetToMeters/F999999999.000000_to_M304799999.695200
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800#01
    length_conversion_test.go:31: Given feet=1.000000, returned meters=0.304800, <nil>, want match for 0.304799970,nil        
=== RUN   TestConvertFeetToMeters/F1.121255_to_M0.341758
=== RUN   TestConvertFeetToMeters/F0.000000_to_M0.000000
--- FAIL: TestConvertFeetToMeters (0.01s)
    --- PASS: TestConvertFeetToMeters/F1.000000_to_M0.304800 (0.00s)
    --- PASS: TestConvertFeetToMeters/F99.000000_to_M30.175200 (0.00s)
    --- PASS: TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000 (0.00s)
    --- PASS: TestConvertFeetToMeters/F999999999.000000_to_M304799999.695200 (0.00s)
    --- FAIL: TestConvertFeetToMeters/F1.000000_to_M0.304800#01 (0.00s)
    --- PASS: TestConvertFeetToMeters/F1.121255_to_M0.341758 (0.00s)
    --- PASS: TestConvertFeetToMeters/F0.000000_to_M0.000000 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/variables 0.055s
```

- We find that the error is in the test, so we make the test case more precise

```go
func TestConvertFeetToMeters(t *testing.T){
    feet_array := []struct{
        feet float64
        meters float64
    } {
        {1, 0.3048},
        {99, 30.1752},
        {100000000, 30480000},
        {999999999, 304799999.6952},
        {0.9999999, 0.304799969520000046951}, // Changed this one to a more precise value using a different calculator
        {1.121254778984587, 0.34175845663450216128},
        {0.00000000000000000000000000000001, 3.048000000000000277e-33},
    }
    
    for _, tc := range feet_array {
        t.Run(fmt.Sprintf("F%f to M%f", tc.feet, tc.meters), func(t *testing.T) {
            meters, err := ConvertFeetToMeters(tc.feet)

            if (err != nil) {
                t.Fatalf("Error in test %v", err)
            }

            if (meters != tc.meters){
                t.Errorf("Given feet=%f, returned meters=%.72f, %v, want match for %.72f,nil", tc.feet, meters, err, tc.meters)
            }
        })
    }
}
```

```powershell
programs/June/variables on  gobook [!?⇡] via 🐹 v1.16.5 on ☁️  (us-east-1)
❯ go test -v -run TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800
=== RUN   TestConvertFeetToMeters/F99.000000_to_M30.175200
=== RUN   TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000
=== RUN   TestConvertFeetToMeters/F999999999.000000_to_M304799999.695200
=== RUN   TestConvertFeetToMeters/F1.000000_to_M0.304800#01
=== RUN   TestConvertFeetToMeters/F1.121255_to_M0.341758
=== RUN   TestConvertFeetToMeters/F0.000000_to_M0.000000
--- PASS: TestConvertFeetToMeters (0.01s)
    --- PASS: TestConvertFeetToMeters/F1.000000_to_M0.304800 (0.00s)
    --- PASS: TestConvertFeetToMeters/F99.000000_to_M30.175200 (0.00s)
    --- PASS: TestConvertFeetToMeters/F100000000.000000_to_M30480000.000000 (0.00s)
    --- PASS: TestConvertFeetToMeters/F999999999.000000_to_M304799999.695200 (0.00s)
    --- PASS: TestConvertFeetToMeters/F1.000000_to_M0.304800#01 (0.00s)
    --- PASS: TestConvertFeetToMeters/F1.121255_to_M0.341758 (0.00s)
    --- PASS: TestConvertFeetToMeters/F0.000000_to_M0.000000 (0.00s)
PASS
ok      debabrata.xyz/variables 0.097s
```

This ends chapter 

## Additional notes

### Language specific

- Golang

```go

```

- Powershell

```powershell

```
