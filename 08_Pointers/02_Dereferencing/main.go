package main

import (
	"fmt"
)

func main() {

	a := 42

	fmt.Println(a)  // 42
	fmt.Println(&a) // 0xc420088008

	var b = &a
	fmt.Println(b)  // 0xc420014058
	fmt.Println(*b) // 42

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
