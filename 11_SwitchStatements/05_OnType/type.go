package main

import (
	"fmt"
)

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

func main() {

	SwitchOnType(7)
	SwitchOnType("Kim")
	var t = contact{"Hello", "Papure"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("String")
	case contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Unknown")
	}
}
