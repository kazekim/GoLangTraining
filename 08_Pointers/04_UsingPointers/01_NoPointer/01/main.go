package main

import (
	"fmt"
)

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}

func zero(z int) {
	z = 0
}
