package main 

import (
	"fmt"
)

func main(){

	var n complex64 = 1 +2i
	fmt.Printf("%v,%T\n", n,n) // output (1+2i),complex64

	fmt.Printf("%v,%T\n", imag(n), imag(n)) //output 2,float32
	fmt.Printf("%v,%T\n", real(n), real(n)) //1,float32
	var x complex128 = complex(5,12)

	fmt.Printf("%v, %v\n", real(x), imag(x)) //output 5,12

}
