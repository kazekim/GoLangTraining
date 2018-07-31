package main

import "fmt"

func main() {
	switch "Kim" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Kim":
		fmt.Println("Wassup Kim")
		fallthrough // if has fallthrough, next case will run too.
	case "Papure":
		fmt.Println("Wassup Papure")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
