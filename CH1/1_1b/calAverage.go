package main

import "fmt"

func main() {
	scores := []int{81, 82, 83, 90, 92}
	sum := 0
	var average float64
	len := len(scores)

	for i := 0; i < len; i++ {
		sum += scores[i]
	}

	average = float64(sum) / float64(len)

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("average: %.2f\n", average)
}
