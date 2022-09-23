package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Please provide at least two numbers")
		return
	}

	nums := []int{}
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err == nil {
			nums = append(nums, num)
		}
	}

	fmt.Println("Before sorting:", nums)
	sort.Ints(nums)
	fmt.Println("After sorting:", nums)
}
