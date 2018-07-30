package main

import (
	"fmt"
)

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

func zero(z *int) {
	*z = 0
}
