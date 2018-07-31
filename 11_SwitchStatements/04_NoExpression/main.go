package main

import (
	"fmt"
)

func main() {

	myFriendName := "Jib"

	switch {
	case len(myFriendName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendName == "Kim":
		fmt.Println("Wassup Kim")
	case myFriendName == "Papure":
		fmt.Println("Wassup Papure")
	case myFriendName == "Jib", myFriendName == "A":
		fmt.Println("Your name is either Jib or A")
	case myFriendName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
