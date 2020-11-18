package main

import (
	"fmt"
)

func main() {
	defer FirstRun()
	SecondRun()

}
func FirstRun()  { 
	
	fmt.Println("I executed First")
	}
func SecondRun() {
	
	fmt.Println("I executed Second") 
	}
