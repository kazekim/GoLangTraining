package main

import (
	"fmt"
)

const (
	_ = iota      // 0
	b = 10 * iota // 1 * 10
	c = 10 * iota // 2 * 10
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
}
