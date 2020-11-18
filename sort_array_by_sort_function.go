package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{5, 36, 26, 1, 2, 698, 6, 65, 133, 65, 4, 2, 32}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

}
