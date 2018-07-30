package main

import (
	"fmt"
)

func main() {

	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}
