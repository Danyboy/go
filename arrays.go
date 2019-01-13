package main

import (
	"fmt"
)

func main2() {
	var a1 [3]int //000

	fmt.Printf("Print %v\n", a1)
	fmt.Printf("Print %#v\n", a1)

	const c1 int = 2
	var a2[2 * c1]bool

	fmt.Printf("Print %#v\n", a2)

	a3 := [...]int{1,2,3}

	fmt.Printf("a3 %#v\n", a3)


}
