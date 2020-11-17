package main

import (
	"fmt"
)

func main() {
	var age int =18
	switch age{
		case 16: fmt.Println("Prepare for college")
		case 18: fmt.Println("Don't smoke")
		case 20: fmt.Println("Get yourself a job")
		default: fmt.Println("Are you alive")
	}
}

