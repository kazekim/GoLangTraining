package main

import (
	"fmt"

	"github.com/gostarter/04_Scope/01_PackageScope/02_Visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
