# An Introduction to Programming in Go: notes for chapters 6

Book by Caleb Doxsey.

<!-- markdownlint-disable MD010 -->

## Sections

- [An Introduction to Programming in Go: notes for chapters 6](#an-introduction-to-programming-in-go-notes-for-chapters-6)
  - [Sections](#sections)
  - [Notes](#notes)
  - [Arrays, Slices and Maps](#arrays-slices-and-maps)
  - [Slices](#slices)
  - [Maps](#maps)
  - [Exercise Programs](#exercise-programs)
    - [What is the length of a slice created using: `make([]int, 3, 9)`?](#what-is-the-length-of-a-slice-created-using-makeint-3-9)
    - [Given the array: `x := [6]string{"a","b","c","d","e","f"}` what would `x[2:5]` give you?](#given-the-array-x--6stringabcdef-what-would-x25-give-you)
    - [Write a program that finds the smallest number in this list](#write-a-program-that-finds-the-smallest-number-in-this-list)
  - [Additional notes](#additional-notes)
    - [Language specific](#language-specific)

## Notes

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

## Maps

- Maps are unordered key-value pair collections.
- Maps can be created as `map[<key type>]<value type>`.

```go
// mappings.go
package main

import "fmt"

func main() {
	var x map[string]int

	x["key"] = 10
	fmt.Println(x)
}

```

```log
[Running] go run "j:\Education\Code\Go\golang-book-notes\books\intro-to-prog-in-go-caleb-doxsey\programs\September\mapping.go"
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
	j:/Education/Code/Go/golang-book-notes/books/intro-to-prog-in-go-caleb-doxsey/programs/September/mapping.go:8 +0x2e
exit status 2

[Done] exited with code=1 in 1.162 seconds
```

- This error happens because the map needs to be initialized before use.

```go
// mapping.go
package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["key"] = 10
	fmt.Println(x)
}

```

```poweshell
intro-to-prog-in-go-caleb-doxsey/programs/September on  gobook [!?] via 🐹 
v1.17 on ☁️  (us-east-1) 
❯ go run .\mapping.go
map[key:10]

intro-to-prog-in-go-caleb-doxsey/programs/September on  gobook [!?] via 🐹 
v1.17 on ☁️  (us-east-1) 
❯ 
```

- We initialize the `map` with `make` function. This creates a `map` object for us to use.
- We can also create maps of `int` keys or any type of keys that are comparable including struct (but not functions and slices) and any type of values.
- We can find the lenght of a map with `len` function.
- We can delete a key and dereference it's value with the `delete(map, key)` function.
- We can also initialize empty maps with `{}`.

```go
// mapping.go
package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["key"] = 10
	fmt.Println(x)

	y := make(map[int]int)
	y[0] = 0
	y[1] = 1
	fmt.Println(y)
	fmt.Println(len(y))
	delete(y, 0)
	fmt.Println(y)
	fmt.Println(len(y))

	z := map[string]float32{}

	z["red"] = 0.01
	fmt.Println(z)
}

```

```powershell
intro-to-prog-in-go-caleb-doxsey/programs/September on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯ go run .\mapping.go
map[key:10]
map[0:0 1:1]
2
map[1:1]
1
map[red:0.01]

intro-to-prog-in-go-caleb-doxsey/programs/September on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- Maps are not sequential.
- If a key doesn't exist in a map, it returns a zero value for the value type. For a string it returns a `""`.
- Accessing an element of a map can return 2 values instead of one.
- The firstvalue is the lookup and the second is if the lookup was successful, a `bool` type.

```go
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

```

```powershell
intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1) 
❯ go run .\mapping2.go
Printing not found int value: 0 , lookup success value: false
Printing not found float value: 0 , lookup success value: false
Printing not found string value:  , lookup success value: false
Lookup successful, 10, true
Here in lookup unsuccessful land

intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- Here we see the default value for `int`, `float32`, and `float64` are 0, while the default value for a string is `''`, or an empty string.
- The if else here works by evaluating the value after the `;`. Here is an example program to demonstrate it.

```go
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
```

```powershell
intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1) 
❯ go run .\if_else_this_way.go
true
false
This prints
This prints again

intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- Values can also be added to maps during creation.

```go
// mapping3.go
package main

import "fmt"

func main (){
    elements := map[string]string {
        "H": "Hydrogen",
        "He": "Helium",
    }
    fmt.Println(elements)
}
```

```powershell
intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯ go run .\mapping3.go
map[H:Hydrogen He:Helium]

intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- We can also store maps as values.

```go
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

```

```powershell

intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯ go run .\mapping4.go
map[H:map[atomic_number:1 name:Hydrogen] He:map[atomic_number:2 name:Helium]]

intro-to-prog-in-go-caleb-doxsey/programs/October on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

## Exercise Programs

### What is the length of a slice created using: `make([]int, 3, 9)`?

- The answer is 3. It creates a slice with 3 0s and the capacity of the undelying array is 9.
- [cap vs len of slice in golang](https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang)

### Given the array: `x := [6]string{"a","b","c","d","e","f"}` what would `x[2:5]` give you?

- The answer is `["c","d","e"]`.

### Write a program that finds the smallest number in this list

```go
x := []int{
48,96,86,68,
57,82,63,70,
37,34,83,27,
19,97, 9,17,
}
```

- We create a new module `smallest` inside the `smallest` folder with `go mod init debabrata.xyz/smallest`

```powershell

October/exercise/smallest on  gobook [!?] on ☁️  (us-east-1) 
❯ go mod init debabrata.xyz/smallest
go: creating new go.mod: module debabrata.xyz/smallest

October/exercise/smallest on  gobook [!?] via 🐹 v1.17.2 on ☁️  (us-east-1) 
❯ 
```

- We create a skeleton file `smallest.go`

```go
// smallest.go

package smallest

import (
	"errors"
	"fmt"
)

var ErrEmptyList = errors.New("list is empty: supply a list with at least one member")

func main(){
    fmt.Println("Printing smallest in the list:")
    fmt.Println("") // TODO: Add a list
    fmt.Printf("Smallest in the list is %v", 1) // TODO: Substitute with actual value computed to be smallest
}

// returns the smallest value in an unsorted slice of ints
func Smallest(numlist []int) (int, error){
    return 1, nil
}
```

- We create a test file with test cases as `smallest_test.go`

```go
// smallest_test.go

package smallest

import (
	"fmt"
	"testing"
)

func TestSmallest(t *testing.T) {
	num_array := []struct {
		nums         []int
		smallest     int
		err_expected error
	}{
		{[]int{1}, 1, nil},
		{[]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2}, -1, nil},
		{[]int{90, 30, 12, 12, 1200}, 12, nil},
		{[]int{4, 1, 1, 0, 2, 8}, 0, nil},
		{[]int{}, 0, ErrEmptyList}, // 0 is a placeholder, can be replaced with whatever, it is never used.
	}

	for _, tc := range num_array {
		t.Run(fmt.Sprintf("N%d", tc.smallest), func(t *testing.T) {
			smallest, err := Smallest(tc.nums)
			if err != tc.err_expected {
				t.Fatalf("Error in test %v. Expected error is %v.", err, tc.err_expected)
			}

			if smallest != tc.smallest {
                t.Errorf("Returned smallest = %d, %v. Want match for %d, %v.", smallest, err, tc.smallest, tc.err_expected)
			}
		})
	}
}

```

- We run the test with `go test -v -run TestSmallest`

```powershell
ctober/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 5s
❯ go test -v -run TestSmallest
=== RUN   TestSmallest
=== RUN   TestSmallest/N1
=== RUN   TestSmallest/N-1
    smallest_test.go:31: Returned smallest = 1, <nil>. Want match for -1, <nil>.
=== RUN   TestSmallest/N12
    smallest_test.go:31: Returned smallest = 1, <nil>. Want match for 12, <nil>.
=== RUN   TestSmallest/N0
    smallest_test.go:31: Returned smallest = 1, <nil>. Want match for 0, <nil>.
=== RUN   TestSmallest/N0#01
    smallest_test.go:27: Error in test <nil>. Expected error is List is empty. Supply a list of ints..
--- FAIL: TestSmallest (0.01s)
    --- PASS: TestSmallest/N1 (0.00s)
    --- FAIL: TestSmallest/N-1 (0.00s)
    --- FAIL: TestSmallest/N12 (0.00s)
    --- FAIL: TestSmallest/N0 (0.00s)
    --- FAIL: TestSmallest/N0#01 (0.00s)
FAIL
exit status 1
FAIL    debabrata.xyz/smallest  0.103s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- As we can see only one test case passes by default.
- We write the logic to find the smallest using a simple linear search.

```go
// returns the smallest value in an unsorted slice of ints
func Smallest(numlist []int) (int, error){
    smallest := numlist[0]
    for _, num := range numlist{
        if num < smallest{
            smallest = num
        }
    }
    return smallest, nil
}
```

- We run the test with `go test -v -run TestSmallest`

```powershell
October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) 
❯ go test -v -run TestSmallest
=== RUN   TestSmallest
=== RUN   TestSmallest/N1
=== RUN   TestSmallest/N-1
=== RUN   TestSmallest/N12
=== RUN   TestSmallest/N0
=== RUN   TestSmallest/N0#01
--- FAIL: TestSmallest (0.00s)        
    --- PASS: TestSmallest/N1 (0.00s) 
    --- PASS: TestSmallest/N-1 (0.00s)
    --- PASS: TestSmallest/N12 (0.00s)
    --- PASS: TestSmallest/N0 (0.00s)
    --- FAIL: TestSmallest/N0#01 (0.00s)
panic: runtime error: index out of range [0] with length 0 [recovered]
        panic: runtime error: index out of range [0] with length 0

goroutine 24 [running]:
testing.tRunner.func1.2({0x9d5de0, 0xc00009e0d8})
        C:/Program Files/Go/src/testing/testing.go:1209 +0x24e
testing.tRunner.func1()
        C:/Program Files/Go/src/testing/testing.go:1212 +0x218
panic({0x9d5de0, 0xc00009e0d8})
        C:/Program Files/Go/src/runtime/panic.go:1038 +0x215
debabrata.xyz/smallest.Smallest({0xb227e8, 0x474, 0x97fba0})
        J:/Education/Code/Go/golang-book-notes/books/intro-to-prog-in-go-caleb-doxsey/programs/October/exercise/smallest/smallest.go:20 +0x4d
debabrata.xyz/smallest.TestSmallest.func1(0xc0000c5d40)
        J:/Education/Code/Go/golang-book-notes/books/intro-to-prog-in-go-caleb-doxsey/programs/October/exercise/smallest/smallest_test.go:25 +0x4a
testing.tRunner(0xc0000c5d40, 0xc0000846b0)
        C:/Program Files/Go/src/testing/testing.go:1259 +0x102
created by testing.(*T).Run
        C:/Program Files/Go/src/testing/testing.go:1306 +0x35a
exit status 2
FAIL    debabrata.xyz/smallest  0.263s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- As we can see that there is an index index out of range error since we passes in an empty list in the last test case.

```go
{[]int{}, 0, ErrEmptyList}, // 0 is a placeholder, can be replaced with whatever, it is never used.
```

- We will thus raise the error we defined in `smallest.go` as `ErrEmptyList`.

```go
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
```

- We run the test with `go test -v -run TestSmallest`

```powershell

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯ go test -v -run TestSmallest
=== RUN   TestSmallest
=== RUN   TestSmallest/N1
=== RUN   TestSmallest/N-1
=== RUN   TestSmallest/N12
=== RUN   TestSmallest/N0
=== RUN   TestSmallest/N0#01
--- PASS: TestSmallest (0.00s)        
    --- PASS: TestSmallest/N1 (0.00s) 
    --- PASS: TestSmallest/N-1 (0.00s)
    --- PASS: TestSmallest/N12 (0.00s)
    --- PASS: TestSmallest/N0 (0.00s)
    --- PASS: TestSmallest/N0#01 (0.00s)
PASS
ok      debabrata.xyz/smallest  0.061s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯
```

- We will try improving performance by sorting the list and then returning the first index, which should make performance better.

- We will first write a benchmark test.

```go
var smallest_result int

func BenchmarkSmallest(b *testing.B){
    var result int
    for n := 0; n <b.N; n++{
        result, _ = Smallest([]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2})
    }
    smallest_result = result
}
```

- We run the benchmark with `go test -bench=BenchmarkSmallest -benchtime=20s` which runs the tests and the benchmark code.

```powershell
October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1)
❯ go test -bench=BenchmarkSmallest -benchtime=20s
goos: windows
goarch: amd64
pkg: debabrata.xyz/smallest
cpu: Intel(R) Core(TM)2 Duo CPU     E7400  @ 2.80GHz
BenchmarkSmallest-2     912021814               27.89 ns/op
PASS
ok      debabrata.xyz/smallest  28.386s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 29s
❯
```

- We write the optimization code.

```go
// returns the smallest value in an unsorted slice of ints
func Smallest(numlist []int) (int, error){
    if len(numlist) == 0 {
        return 0, ErrEmptyList
    }
    sort.Ints(numlist)
    
    return numlist[0], nil
}
```

- We run the benchmark with `go test -bench=BenchmarkSmallest -benchtime=20s` which runs the tests and the benchmark code.

```powershell
October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 29s
❯ go test -bench=BenchmarkSmallest -benchtime=20s
goos: windows
goarch: amd64
pkg: debabrata.xyz/smallest
cpu: Intel(R) Core(TM)2 Duo CPU     E7400  @ 2.80GHz
BenchmarkSmallest-2     46377528               511.4 ns/op
PASS
ok      debabrata.xyz/smallest  44.062s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 45s
❯
```

- As we can see the benchmark shows the code is prematurely optimized. We roll back the optimizations.

```go
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
```

- We run the benchmark again to be sure.

```powershell
October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 23s
❯ go test -bench=BenchmarkSmallest -benchtime=20s
goos: windows
goarch: amd64
pkg: debabrata.xyz/smallest
cpu: Intel(R) Core(TM)2 Duo CPU     E7400  @ 2.80GHz
BenchmarkSmallest-2     902104959               26.58 ns/op
PASS
ok      debabrata.xyz/smallest  26.827s

October/exercise/smallest on  gobook [✘?] via 🐹 v1.17.2 on ☁️  (us-east-1) took 28s
❯
```

- We cleanup the code.

```go
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
```

```go
// smallest_test.go

package smallest

import (
	"fmt"
	"testing"
)

func TestSmallest(t *testing.T) {
	num_array := []struct {
		nums         []int
		smallest     int
		err_expected error
	}{
		{[]int{1}, 1, nil},
		{[]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2}, -1, nil},
		{[]int{90, 30, 12, 12, 1200}, 12, nil},
		{[]int{4, 1, 1, 0, 2, 8}, 0, nil},
		{[]int{}, 0, ErrEmptyList}, // 0 is a placeholder, can be replaced with whatever, it is never used.
	}

	for _, tc := range num_array {
		t.Run(fmt.Sprintf("N%d", tc.smallest), func(t *testing.T) {
			smallest, err := Smallest(tc.nums)
			if err != tc.err_expected {
				t.Fatalf("Error in test %v. Expected error is %v.", err, tc.err_expected)
			}

			if smallest != tc.smallest {
                t.Errorf("Returned smallest = %d, %v. Want match for %d, %v.", smallest, err, tc.smallest, tc.err_expected)
			}
		})
	}
}

var smallest_result int

func BenchmarkSmallest(b *testing.B){
    var result int
    for n := 0; n <b.N; n++{
        result, _ = Smallest([]int{20, 58, 463, 65, 22222, -1, 88, 554475, 66352, 2})
    }
    smallest_result = result
}

```

## Additional notes

### Language specific

- Golang

```go

```

- Powershell

```powershell

```
