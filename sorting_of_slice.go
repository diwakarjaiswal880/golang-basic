package main 
  
import ( 
    "fmt"
    "sort"
) 
  
func main() { 
  
    // Creating Slice 
    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"} 
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89} 
  
    fmt.Println("Before sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
  
    // Performing sort operation on the 
    // slice using sort function 
    sort.Strings(slc1) 
    sort.Ints(slc2) 
  
    fmt.Println("\nAfter sorting:") 
    fmt.Println("Slice 1: ", slc1) 
    fmt.Println("Slice 2: ", slc2) 
  
} 
