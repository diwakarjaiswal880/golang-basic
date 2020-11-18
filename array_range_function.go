package main

import "fmt"

func main(){
	student:=[5]string {"ram","mohan","geta"}
	for i,value := range(student){
		fmt.Println(i,value)
	
	}
	student[4]="Hi"
	fmt.Println(student[4],"mohan")

}
