package main 
  
import "fmt"
  
func main() { 
  
    // Creating multi-dimensional slice 
    s1 := [][]int{{12, 34}, 
        {56, 47}, 
        {29, 40}, 
        {46, 78}, 
    } 
  
    // Accessing multi-dimensional slice 
    fmt.Println("Slice 1 : ", s1) 
  
    // Creating multi-dimensional slice 
    s2 := [][]string{ 
        []string{"Geeks", "for"}, 
        []string{"Geeks", "GFG"}, 
        []string{"gfg", "geek"}, 
    } 
  
    // Accessing multi-dimensional slice 
    fmt.Println("Slice 2 : ", s2) 
  
}
