package main

import "fmt"

func main() {
	const (
		min = 1
		max = 10
	)
	sum := 0

	for i := min; i <= max; i++ {
		fmt.Printf("%d", i)
		if i < max {
			fmt.Printf(" + ")
		}
		sum += i
	}

	fmt.Printf("= %d\n", sum)
}
