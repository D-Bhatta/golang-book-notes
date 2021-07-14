# An Introduction to Programming in Go: notes for chapters 1-4

Book by Caleb Doxsey.

<!-- markdownlint-disable MD010 -->

## Sections

- [An Introduction to Programming in Go: notes for chapters 1-4](#an-introduction-to-programming-in-go-notes-for-chapters-1-4)
  - [Sections](#sections)
  - [Notes](#notes)
  - [Control Structures: For Loop](#control-structures-for-loop)
  - [Additional notes](#additional-notes)
    - [Language specific](#language-specific)

## Notes

## Control Structures: For Loop

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

- Running the program with `go run  .\for_loop.go` prints numbers from 1 till 10.

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

- This is another way of using the for loop. Running it with `go run  .\for_loop.go` gives us the following output.

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



## Additional notes

### Language specific

- Golang

```go

```

- Powershell

```powershell

```
