package main

import (
	"fmt"
)

func main() {
	greet("Kim", "Papure")
}

func greet(name1, name2 string) {
	fmt.Println(name1, " - ", name2)
}
