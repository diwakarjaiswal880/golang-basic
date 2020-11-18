package main 
  
import "fmt"
  
func main() { 
  
    // Creating a zero value slice 
    arr := [6]int{55, 66, 77, 88, 99, 22} 
    slc := arr[0:4] 
  
    // Before modifying 
  
    fmt.Println("Original_Array: ", arr) 
    fmt.Println("Original_Slice: ", slc) 
  
    // After modification 
    slc[0] = 100 
    slc[1] = 1000 
    slc[2] = 1000 
  
    fmt.Println("\nNew_Array: ", arr) 
    fmt.Println("New_Slice: ", slc) 
}
