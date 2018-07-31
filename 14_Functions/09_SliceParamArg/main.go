package main

import (
	"fmt"
)

func main() {
	data := []float64{18, 19, 30, 44, 60, 72}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
