package main

import "fmt"

// this variable is initialized first due to
// order of declaration
var initCounter int

func init() {
	fmt.Println("Called First in Order of Declaration")
	initCounter++
}

func init() {
	fmt.Println("Called second in order of declaration")
	initCounter++
}

func main() {
	fmt.Println("Does nothing of any significance")
	fmt.Printf("Init Counter: %d\n", initCounter)
}
