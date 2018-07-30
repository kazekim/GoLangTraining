package main

import "fmt"

func main() {
	max := max(7)
	fmt.Println(max)
}

func max(x int) int {
	return x + 42
}
