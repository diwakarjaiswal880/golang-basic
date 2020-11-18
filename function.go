package main 
  
import ( 
    "fmt"
) 
  
func main() { 
	x,y:=5,6
	fmt.Println(add(x,y))

}
func add(num1,num2 int) int{
	return num1+num2
} 
