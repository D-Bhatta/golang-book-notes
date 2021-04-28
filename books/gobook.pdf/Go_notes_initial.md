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




## Additional notes

### Language specific

- Golang

```go

```

- Powershell

```powershell

```
