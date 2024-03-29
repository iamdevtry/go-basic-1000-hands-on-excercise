package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		nums   []int
		oldCap float64
	)
	start := time.Now()
	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
