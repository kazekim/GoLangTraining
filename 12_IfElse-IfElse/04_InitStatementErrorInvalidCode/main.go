package main

import (
	"fmt"
)

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	//can't find variable "food" here
	fmt.Println(food)
}
