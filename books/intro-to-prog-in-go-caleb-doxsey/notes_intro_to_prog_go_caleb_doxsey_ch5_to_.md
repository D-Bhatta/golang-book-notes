# An Introduction to Programming in Go: notes for chapters 1-4

Book by Caleb Doxsey.

<!-- markdownlint-disable MD010 -->

## Sections

- [An Introduction to Programming in Go: notes for chapters 1-4](#an-introduction-to-programming-in-go-notes-for-chapters-1-4)
  - [Sections](#sections)
  - [Notes](#notes)
  - [Control Structures](#control-structures)
    - [For Loop](#for-loop)
    - [If Statement](#if-statement)
    - [Switch statement](#switch-statement)
  - [Exercise Programs: Chapter 5](#exercise-programs-chapter-5)
    - [All numbers evenly divisible by 3 between 1 and 100](#all-numbers-evenly-divisible-by-3-between-1-and-100)
    - [Fizz buzz 1 to 100](#fizz-buzz-1-to-100)
  - [Arrays, Slices and Maps](#arrays-slices-and-maps)
  - [Slices](#slices)
  - [Additional notes](#additional-notes)
    - [Language specific](#language-specific)

## Notes

## Control Structures

### For Loop

- The for loop evaluates an expression, and if the expression is `true`, it executes a block of code. This repeats until the expression evaluates to false.

```go
//  for_loop.go
package main

import "fmt"

func main(){
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }
}
```

- Running the program with `go run .\for_loop.go` prints numbers from 1 till 10.

```powershell
programs/July/control_structures on  gobook [!?] via 🐹 v1.16.5 on ☁️  (us-east-1) took 4s
❯ go run  .\for_loop.go
1
2
3
4
5
6
7
8
9
10

programs/July/control_structures on  gobook [!?] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯
```

- Go only has one loop.

```go
//  for_loop.go
package main

import "fmt"

func main(){
    // i := 1
    // for i <= 10 {
    //     fmt.Println(i)
    //     i = i + 1
    // }
    for j := -5; j <= 0; j++ {
        fmt.Println(j)
    }
}
```

- This is another way of using the for loop. Running it with `go run .\for_loop.go` gives us the following output.

```powershell
programs/July/control_structures on  gobook [!?] via 🐹 v1.16.5 on ☁️  (us-east-1) took 2s
❯ go run  .\for_loop.go
-5
-4
-3
-2
-1
0

```

### If Statement

- In Go, the `%` returns the remainder after division.
- To create an even/odd program, we use the `if` statement to detemine if a statement is even or odd.
- `else` block executes when `if` condition evaluates to false.
- There can be `else if` parts, which can be used to construct a logic ladder.

```go
package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d is odd\n", i)
		} else {
			fmt.Println("This statement never prints")
		}
	}
}

```

### Switch statement

- Another way to construct a logic ladder is the `switch` statement.
- It uses the case keyword to evaluate the expression against each possibility.

```go
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

```

- It takes a number as input, evaluates it against each possible case and executes the case.
- If no case can be right, the `default` keyword is used to catch the base case.

```powershell
programs/July/control_structures on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 9s
❯ go run .\switch_statement.go
Choses case 1

programs/July/control_structures on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 3s
❯ go run .\switch_statement.go
Enter a number between 1 to 4: 2
Choses case 2

programs/July/control_structures on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 2s
❯ go run .\switch_statement.go
Enter a number between 1 to 4: 3
Choses case 3

programs/July/control_structures on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1)
❯ go run .\switch_statement.go
Enter a number between 1 to 4: 4
Choses case 4

programs/July/control_structures on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 2s
❯ go run .\switch_statement.go
Enter a number between 1 to 4: 15
Chosen number is out of range
```

## Exercise Programs: Chapter 5

First we create a module `go mod init debabrata.xyz/ctrl_structures`.

### All numbers evenly divisible by 3 between 1 and 100

- We need to find all numbers between 1 and 100 which can be divided by 3 and leave 0 as the remainder.
- We shall use the following test cases to start work:

| Number | Remainder |
| :----- | --------: |
| 1      |         1 |
| 3      |         0 |
| 6      |         0 |
| 0      |         0 |
| 9      |         0 |
| 12     |         0 |

- We create a skeleton function

```go
package ctrlstructs

import (
	"fmt"
)

func DivBy3Printer() {
	numbers := 3
	fmt.Print(numbers)

}

/* DivBy3Util returns an array of numbers that are divisible by 3 from 0 till end. */
func DivBy3Util(end int) ([]int, error) {
	var numbers []int
	return numbers, nil
}
```

- We create a test file with one test case

```go
package ctrlstructs

import (
	"fmt"
	"testing"
)

func TestDivBy3Util(t *testing.T) {
	numbers_array := []struct {
		end     int
		numbers []int
	}{
		{0, []int{0}},
		{10, []int{0, 3, 6, 9}},
		{100, []int{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}},
	}

	for _, tc := range numbers_array {
		t.Run(fmt.Sprintf("E%d", tc.end), func(t *testing.T) {
			numbers, err := DivBy3Util(tc.end)

			if err != nil {
				t.Fatalf("Error in test %v", err)
			}

			if len((numbers)) != len(tc.numbers) {
				t.Errorf("Given end=%d, returned len(numbers)=%d, %v, want match for len(numbers)=%d,nil", tc.end, len(numbers), err, len(tc.numbers))
			} else {
				for i := 0; i < len(tc.numbers); i++ {
					if numbers[i] != tc.numbers[i] {
						t.Errorf("Given end=%d, returned numbers[%d]=%d, %v, want match for numbers[%d]=%d,nil", tc.end, i, numbers[i], err, i, tc.numbers[i])
					}
				}
			}

		})
	}
}


```

- We run the test with `go test -v -run TestDivBy3Util` and make sure it fails

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 3s
❯ go test -v -run TestDivBy3Util
=== RUN   TestDivBy3Util
=== RUN   TestDivBy3Util/E0
    div_by_3_test.go:27: Given end=0, returned len(numbers)=0, <nil>, want match for len(numbers)=1,nil
=== RUN   TestDivBy3Util/E10
    div_by_3_test.go:27: Given end=10, returned len(numbers)=0, <nil>, want match for len(numbers)=4,nil
=== RUN   TestDivBy3Util/E100
    div_by_3_test.go:27: Given end=100, returned len(numbers)=0, <nil>, want match for len(numbers)=34,nil
--- FAIL: TestDivBy3Util (0.01s)
    --- FAIL: TestDivBy3Util/E0 (0.00s)
    --- FAIL: TestDivBy3Util/E10 (0.00s)
    --- FAIL: TestDivBy3Util/E100 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/ctrlstructs       0.058s
```

- We write the code for the program

```go
package ctrlstructs

import (
	"errors"
	"fmt"
)

func DivBy3Printer() {
	numbers := 3
	fmt.Print(numbers)

}

/* DivBy3Util returns an array of numbers that are divisible by 3 from 0 till end. */
func DivBy3Util(end int) ([]int, error) {
	var numbers []int

	for i := 0; i <= end; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}

	}

	return numbers, nil
}

```

- We run the test with `go test -v -run TestDivBy3Util`

```powershell
❯ go test -v -run TestDivBy3Util
=== RUN   TestDivBy3Util
=== RUN   TestDivBy3Util/E0
=== RUN   TestDivBy3Util/E10
=== RUN   TestDivBy3Util/E100
--- PASS: TestDivBy3Util (0.00s)
    --- PASS: TestDivBy3Util/E0 (0.00s)
    --- PASS: TestDivBy3Util/E10 (0.00s)
    --- PASS: TestDivBy3Util/E100 (0.00s)
PASS
ok      debabrata.xyz/ctrlstructs       0.053s
```

- We handle the error case of input `end` being less than 0
- We document this in the doc string

```go
/* DivBy3Util returns an array of numbers that are divisible by 3 from 0 till end. 
If the end variable is less than 0, return an error
*/
func DivBy3Util(end int) ([]int, error) {
	if end < 0 {
		return []int{}, errors.New("passed parameter end cannot be a negative value")
	}

	var numbers []int

	for i := 0; i <= end; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}

	}

	return numbers, nil
}
```

- We spruce up the function

```go
package ctrlstructs

import (
	"errors"
	"fmt"
	"log"
)

func DivBy3Printer() {
	numbers, err := DivBy3Util(100)
	if err != nil {
        log.Fatal(fmt.Sprintf("Error in function DivBy3Printer in ctrlstructs/div_by_3.go: %v", err))
	}
	fmt.Print(numbers)

}

/* DivBy3Util returns an array of numbers that are divisible by 3 from 0 till end.
If the end variable is less than 0, return an error
*/
func DivBy3Util(end int) ([]int, error) {
	if end < 0 {
		return []int{}, errors.New("passed parameter end cannot be a negative value")
	}

	var numbers []int

	for i := 0; i <= end; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}

	}

	return numbers, nil
}
```

### Fizz buzz 1 to 100

Print numbers from 1 to 100, on every multiple of 3 print "Fizz" instead, on every multiple of 5 print "Buzz" instead, and on every multiple of 3 and 5 print "FizzBuzz" instead.

- We create a skeleton file in `fizz_buzz.go`

```go
// fizz_buzz.go
package ctrlstructs

import (
    "fmt"
)

func PrintFizzBuzz(){
    // TODO: Call FizzBuzz and print the data
    fmt.Print("")
}

/* Returns a string slice from 1 to n of FizzBuzz */
func FizzBuzzUtil(n int) ([]string, error){
    return []string{},nil
}
```

- We create a skeleton test file in `fizz_buzz_test.go`

```go
// fizz_buzz_test.go
package ctrlstructs

import (
	"fmt"
	"testing"
)

func TestFizzBuzzUtil(t *testing.T) {
	fizz_array := []struct {
		n         int
		fizz_buzz []string
	}{
		{1, []string{"1"}},
	}

    for _, tc := range fizz_array {
        t.Run(fmt.Sprintf("N%d", tc.n), func(t *testing.T) {
            fizz_buzz, err := FizzBuzzUtil(tc.n)
            
            if err != nil {
                t.Fatalf("Error in test %v", err)
            }
            
            if len(fizz_buzz) != len(tc.fizz_buzz) {
                t.Fatalf("Given n=%d, returned len(fizz_buzz)=%d, %v, want match for len(fizz_buzz)=%d, nil", tc.n, len(fizz_buzz),err, len(tc.fizz_buzz))
            } else {
                for i:= 0; i < len(fizz_buzz); i++{
                    if fizz_buzz[i] != tc.fizz_buzz[i] {
                        t.Errorf("Given n=%d, returned fizz_buzz[%d]=%s, %v, want match for fizz_buzz[%d]=%s, nil", tc.n, i, fizz_buzz[i], err, i, tc.fizz_buzz[i])
                    }
                }
            }
        })
    }
}
```

- We populate the test cases

```go
	fizz_array := []struct {
		n         int
		fizz_buzz []string
        err_expected error
	}{
		{1, []string{"1"}, nil},
		{10, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}, nil},
		{100, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz", "Fizz", "37", "38", "Fizz", "Buzz", "41", "Fizz", "43", "44", "FizzBuzz", "46", "47", "Fizz", "49", "Buzz", "Fizz", "52", "53", "Fizz", "Buzz", "56", "Fizz", "58", "59", "FizzBuzz", "61", "62", "Fizz", "64", "Buzz", "Fizz", "67", "68", "Fizz", "Buzz", "71", "Fizz", "73", "74", "FizzBuzz", "76", "77", "Fizz", "79", "Buzz", "Fizz", "82", "83", "Fizz", "Buzz", "86", "Fizz", "88", "89", "FizzBuzz", "91", "92", "Fizz", "94", "Buzz", "Fizz", "97", "98", "Fizz", "Buzz"}, nil},
        {-1, []string{}, ErrInvalidInput},
        {0, []string{}, nil},
	}
```

- We run the test with `go test -v -run TestFizzBuzzUtil`

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 18s
❯ go test -v -run TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil/N1
    fizz_buzz_test.go:30: Given n=1, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=1, nil
=== RUN   TestFizzBuzzUtil/N10
    fizz_buzz_test.go:30: Given n=10, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=10, nil
=== RUN   TestFizzBuzzUtil/N100
    fizz_buzz_test.go:30: Given n=100, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=100, nil
=== RUN   TestFizzBuzzUtil/N-1
    fizz_buzz_test.go:26: Error in test <nil>
--- FAIL: TestFizzBuzzUtil (0.01s)
    --- FAIL: TestFizzBuzzUtil/N1 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N10 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N100 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N-1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N0 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/ctrlstructs       0.070s

July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1)
❯
```

- We write code for the error case first

```go
/* Returns a string slice from 1 to n of FizzBuzz */
func FizzBuzzUtil(n int) ([]string, error) {
    if n < 0 {
        return []string{}, ErrInvalidInput
    }
	return []string{}, nil
}
```

- We run the test with `go test -v -run TestFizzBuzzUtil`

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 5s
❯ go test -v -run TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil/N1
    fizz_buzz_test.go:30: Given n=1, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=1, nil        
=== RUN   TestFizzBuzzUtil/N10
    fizz_buzz_test.go:30: Given n=10, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=10, nil      
=== RUN   TestFizzBuzzUtil/N100
    fizz_buzz_test.go:30: Given n=100, returned len(fizz_buzz)=0, <nil>, want match for len(fizz_buzz)=100, nil    
=== RUN   TestFizzBuzzUtil/N-1
--- FAIL: TestFizzBuzzUtil (0.01s)
    --- FAIL: TestFizzBuzzUtil/N1 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N10 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N100 (0.00s)
    --- PASS: TestFizzBuzzUtil/N-1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N0 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/ctrlstructs       0.325s

July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 3s
❯
```

- We write code for the `N1` test case

```go
/* Returns a string slice from 1 to n of FizzBuzz */
func FizzBuzzUtil(n int) ([]string, error) {
	if n < 0 {
		return []string{}, ErrInvalidInput
	}

	fizz_buzz := []string{}

	for i := 1; i <= n; i++ {
		fizz_buzz = append(fizz_buzz, strconv.Itoa(i))
	}

	return fizz_buzz, nil
}
```

- We run the test with `go test -v -run TestFizzBuzzUtil`

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 3s
❯ go test -v -run TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil/N1
=== RUN   TestFizzBuzzUtil/N10
    fizz_buzz_test.go:34: Given n=10, returned fizz_buzz[2]=3, <nil>, want match for fizz_buzz[2]=Fizz, nil        
    fizz_buzz_test.go:34: Given n=10, returned fizz_buzz[4]=5, <nil>, want match for fizz_buzz[4]=Buzz, nil        
    ...       
=== RUN   TestFizzBuzzUtil/N100
    fizz_buzz_test.go:34: Given n=100, returned fizz_buzz[2]=3, <nil>, want match for fizz_buzz[2]=Fizz, nil       
    fizz_buzz_test.go:34: Given n=100, returned fizz_buzz[4]=5, <nil>, want match for fizz_buzz[4]=Buzz, nil
    ...
=== RUN   TestFizzBuzzUtil/N-1
--- FAIL: TestFizzBuzzUtil (0.15s)
    --- PASS: TestFizzBuzzUtil/N1 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N10 (0.02s)
    --- FAIL: TestFizzBuzzUtil/N100 (0.12s)
    --- PASS: TestFizzBuzzUtil/N-1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N0 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/ctrlstructs       2.366s

July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 7s
❯
```

- We write code for the `N10` case

```go
/* Returns a string slice from 1 to n of FizzBuzz */
/* n must be 0 or greater */
func FizzBuzzUtil(n int) ([]string, error) {
	if n < 0 {
		return []string{}, ErrInvalidInput
	}

	fizz_buzz := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fizz_buzz = append(fizz_buzz, "Fizz")
			continue
		}
		if i%5 == 0 {
			fizz_buzz = append(fizz_buzz, "Buzz")
			continue
		}
		fizz_buzz = append(fizz_buzz, strconv.Itoa(i))
	}

	return fizz_buzz, nil
}
```

- We run the test with `go test -v -run TestFizzBuzzUtil`

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 14s 
❯ go test -v -run TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil/N1
=== RUN   TestFizzBuzzUtil/N10
=== RUN   TestFizzBuzzUtil/N100
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[14]=Fizz, <nil>, want match for fizz_buzz[14]=FizzBuzz, nil
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[29]=Fizz, <nil>, want match for fizz_buzz[29]=FizzBuzz, nil
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[44]=Fizz, <nil>, want match for fizz_buzz[44]=FizzBuzz, nil
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[59]=Fizz, <nil>, want match for fizz_buzz[59]=FizzBuzz, nil
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[74]=Fizz, <nil>, want match for fizz_buzz[74]=FizzBuzz, nil
    fizz_buzz_test.go:35: Given n=100, returned fizz_buzz[89]=Fizz, <nil>, want match for fizz_buzz[89]=FizzBuzz, nil
=== RUN   TestFizzBuzzUtil/N-1
=== RUN   TestFizzBuzzUtil/N0
--- FAIL: TestFizzBuzzUtil (0.07s)
    --- PASS: TestFizzBuzzUtil/N1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N10 (0.00s)
    --- FAIL: TestFizzBuzzUtil/N100 (0.06s)
    --- PASS: TestFizzBuzzUtil/N-1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N0 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/ctrlstructs       0.405s

July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 4s
❯
```

- We write the code for the `N100` test case

```go
// Returns a string slice from 1 to n of FizzBuzz
// n must be 0 or greater
func FizzBuzzUtil(n int) ([]string, error) {
	if n < 0 {
		return []string{}, ErrInvalidInput
	}

	fizz_buzz := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizz_buzz = append(fizz_buzz, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fizz_buzz = append(fizz_buzz, "Fizz")
			continue
		}
		if i%5 == 0 {
			fizz_buzz = append(fizz_buzz, "Buzz")
			continue
		}
		fizz_buzz = append(fizz_buzz, strconv.Itoa(i))
	}

	return fizz_buzz, nil
}
```

- We run the test with `go test -v -run TestFizzBuzzUtil`

```powershell
July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 4s
❯ go test -v -run TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil
=== RUN   TestFizzBuzzUtil/N1
=== RUN   TestFizzBuzzUtil/N10
=== RUN   TestFizzBuzzUtil/N100
=== RUN   TestFizzBuzzUtil/N-1
=== RUN   TestFizzBuzzUtil/N0
--- PASS: TestFizzBuzzUtil (0.01s)
    --- PASS: TestFizzBuzzUtil/N1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N10 (0.00s)
    --- PASS: TestFizzBuzzUtil/N100 (0.00s)
    --- PASS: TestFizzBuzzUtil/N-1 (0.00s)
    --- PASS: TestFizzBuzzUtil/N0 (0.00s)
PASS
ok      debabrata.xyz/ctrlstructs       0.498s

July/control_structures/ctrlstructs on  gobook [!?⇡] via 🐹 v1.16.6 on ☁️  (us-east-1) took 4s
❯
```

- The good thing about TDD is that you can focus on the code one test case at a time.
- This greatly reduces the cognitive load.

## Arrays, Slices and Maps

- Arrays are numbered sequence data types.
- They have a fixed length, designated during compilation.
- They have a fixed type.
- Arrays can be created as `<var-name> [<length>]<type>`

```go
// array_1.go
package main

import "fmt"

func main() {
	var numbers [5]float64

	numbers[0] = 98

	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println(total / float64(len(numbers)))

    var other_total float64= 0
    for _,value := range numbers {
        other_total += value
    }
    fmt.Println(other_total/float64(len(numbers)))

    second_array := [5]float64{ 98, 0, 0, 10, 11}
    fmt.Println(second_array)

    third_array := [3]string{
        "hello",
        "I",
        "am",
    }

    fmt.Println(third_array)
}

```

- Running the above example program:

```powershell
gobook.pdf/programs/august on  gobook [!?⇡] via 🐹 v1.17 on ☁️  (us-east-1)
❯ go run .\array_1.go
19.6
19.6
[98 0 0 10 11]
[hello I am]

gobook.pdf/programs/august on  gobook [!?⇡] via 🐹 v1.17 on ☁️  (us-east-1)
❯
```

- We can get the length of an array with the `len` function, just like python.
- The `range` keyword enumerates over the entire range of the array and returns a position and a value for each element.
- We can use `_` to tell the compiler that we don't need a particular return type.
- We can also create arrays using the `<var-name> := [<length>]<type>{<value>, <value>}`
- They can also be broken up by comma.

## Slices

- A slice is a segment of an array.
- They have length and elements can be randomly accessed using `[]`.
- Their length can be changed.
- Length of a slice is 0 at the time of creation.

```go
// slice_1.go
package main

import "fmt"

func main() {
	var first_slice []float64
	fmt.Println("first_slice", first_slice)
	second_slice := make([]float64, 5, 10)
	fmt.Println("second_slice", second_slice)
	array_1 := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("array_1", array_1)
	third_slice := array_1[0:5]
	fmt.Println("third_slice", third_slice)
	fourth_slice := []int{1, 2, 3, 4, 5}
	fmt.Println("fourth_slice", fourth_slice)
	fifth_slice := append(fourth_slice, 98, 99)
	fmt.Println("fifth_slice", fifth_slice)
	sixth_slice := make([]int, 2)
	copy(sixth_slice, fourth_slice)
	fmt.Println("sixth_slice", sixth_slice)
	seventh_slice := make([]int, 20)
	copy(seventh_slice, fourth_slice)
	fmt.Println("seventh_slice", seventh_slice)

}

```

- Running the following we see:

```powershell
gobook.pdf/programs/august on  gobook [!?] via 🐹 v1.17 on ☁️  
(us-east-1) 
❯ go run .\slice_1.go
first_slice []
second_slice [0 0 0 0 0]
array_1 [1 2 3 4 5 6 7 8 9 10]
third_slice [1 2 3 4 5]
fourth_slice [1 2 3 4 5]
fifth_slice [1 2 3 4 5 98 99]
sixth_slice [1 2]
seventh_slice [1 2 3 4 5 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

gobook.pdf/programs/august on  gobook [!?] via 🐹 v1.17 on ☁️  
(us-east-1) 
❯
```

- They can be created as `var <var-name> []<type>`.
- They can be created as `<var-name> := []<type>{element, element,...}`
- They can be created using the `make` function as `<var-name> := make([]<type>, <length>, <capacity>)`.
- Slices are always associated with an array, and using the `make` function, the lenght of the slice will be the specified `length` parameter, and the capacity of the array will be the specified `capacity` parameter.
- Slices can be created as `<var-name> := <array-name>[<low>:<high]`.
- The slice will start at the index `low` and include elements until it reaches `high` element index, which would be elements until the `high-1` index.
- `high` and `low` can both be omitted.
- The `append` function creates a new slice by taking an existing slice as the first argument and appending all the other arguments to it.
- The `copy` function takes a slice as the first argument, and copies elements of a slice passed as the second argument, until either the first slice is filled up, or the second slice has no more elements to copy.

## Additional notes

### Language specific

- Golang

```go

```

- Powershell

```powershell

```