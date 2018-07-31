package main

import "fmt"

func main() {
	switch "TonHor" {
	case "Kim", "Papure":
		fmt.Println("Wassup Kim, or, err, Papure")
	case "TonHor", "Tot":
		fmt.Println("Both of your names start with T")
	case "Jib", "A":
		fmt.Println("You are soul mate.")
	}
}
