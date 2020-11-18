package main

import "fmt"

func main(){

	var student[10] int
	class :=[10]int{1,2,3,4,5,6,7,8,9,10}
  
	for i:=0;i<10;i++{
  
		student[i]=i+1
		fmt.Println(student[i])
    
	}
	for i:=0;i<10;i++{
		fmt.Println(class[i])
    
	}
}
