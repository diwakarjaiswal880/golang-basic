package main

import (
	"fmt"
)

func main() {
	var name string = "Diwakar"
	const pi float64 = 3.16412732
	x := 5
	isbool := true
	
	fmt.Println(len(name))
	fmt.Println(name+" is chill dude.")		
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", isbool)
	fmt.Printf("%t \n", isbool)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 67)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)

}

