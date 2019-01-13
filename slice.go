package main

import (
	"fmt"
)

func main() {
	var b1 = []int{42}

	b2 := make([]int, 5, 10)
	
	fmt.Println(b1, b2)

	fmt.Println(b2[4])
	// fmt.Println(b2[5])

	b2 = append(b2, 9)
	b2 = append(b2, b1...)

	fmt.Printf("b1 %#v\n", b1)	
}
