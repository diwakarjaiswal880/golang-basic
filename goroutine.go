package main 
  
import ( 
    "fmt"
    "time"
) 
  
func display(str string) { 
    for w := 0; w < 6; w++ { 
        time.Sleep(1 * time.Second) 
        fmt.Println(str) 
    } 
} 
  
func main() { 
  
    // Calling Goroutine 
    go display("Welcome") 
  
    // Calling normal function 
    display("GeeksforGeeks") 
} 
