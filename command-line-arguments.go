package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

// For run go program through command line arguments
// First build program
//$ go build command-line-arguments.go
// Then run with command line arguments
//$ ./command-line-arguments a b c d
