package main
import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/Hello",handler2)
	http.ListenAndServe(":8081",nil)
}
func handler(w http.ResponseWriter,r * http.Request ) {
	fmt.Fprintf(w, "Welcome to my page. \n")

}
func handler2(w http.ResponseWriter,r * http.Request ) {
	fmt.Fprintf(w, "Hello World. \n")
	
}
