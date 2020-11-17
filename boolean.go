package main

import (
	"fmt"
)

func main() {
	x,y:=5,6		
	fmt.Println("y + y = ",x+y)
	fmt.Println("y - y = ",x-y)
	fmt.Println("y * y = ",x*y)
	fmt.Println("y / y = ",x/y)
	fmt.Println("x mod y = ",x%y)
	
	isbool :=true
	// var isbool bool= true
	hate :=false
	
	fmt.Println("isbool && hate = ", isbool && hate)
	fmt.Println("isbool || hate = ", isbool || hate)
	fmt.Println("! isbool = ", !isbool)

	

}
