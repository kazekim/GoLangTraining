package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Print("Input larger number : ")
	fmt.Scan(&num1)
	fmt.Print("Input smaller number : ")
	fmt.Scan(&num2)

	fmt.Println(num1, " % ", num2, " = ", num1%num2)

}
