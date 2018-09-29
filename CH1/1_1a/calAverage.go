package main

import "fmt"

func main() {
	g1, g2, g3, g4, g5 := 80, 81, 90, 92, 100
	var sum int
	var average float64
	sum = g1 + g2 + g3 + g4 + g5
	average = float64(sum) / 5

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("average: %f\n", average)
}
