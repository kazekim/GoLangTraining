package main

import (
	"fmt"
)

func main() {

	greet("Kim")
	greet("Papure")
}

func greet(name string) {
	fmt.Println("Hello, ", name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
