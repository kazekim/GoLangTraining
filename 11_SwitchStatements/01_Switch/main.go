package main

import (
	"fmt"
)

func main() {
	// no need for "break" in case
	switch "Kim" {
	case "Daniel":
		fmt.Println("Wassap Daniel.")
	case "Kim":
		fmt.Println("Wassap Kim.")
	case "Papure":
		fmt.Println("Wassap Papure.")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
