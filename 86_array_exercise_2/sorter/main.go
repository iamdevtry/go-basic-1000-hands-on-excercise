package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	nums := [5]float64{}

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		nums[i] = n
	}

	fmt.Printf("%v default array\n", nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Printf("%v sorted array\n", nums)
}
