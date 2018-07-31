package main

import (
	"fmt"
)

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	if food := "Chocolate"; !b {
		fmt.Println(food)
	}
}
