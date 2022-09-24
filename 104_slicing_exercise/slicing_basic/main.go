package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"

	splitted := strings.Fields(data)

	fmt.Println(splitted)

	var nums []int
	for _, v := range splitted {
		n, err := strconv.Atoi(v)
		if err == nil {
			nums = append(nums, n)
		}
	}
	fmt.Println("nums: ", nums)

	evens, odds := nums[:3], nums[3:]
	fmt.Println("evens: ", evens)
	fmt.Println("odds: ", odds)

	middle, firstTwo, lastTwo := nums[2:4], nums[:2], nums[len(nums)-2:]
	fmt.Println("middle: ", middle)
	fmt.Println("firstTwo: ", firstTwo)
	fmt.Println("lastTwo: ", lastTwo)

	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2 :", odds[len(odds)-2:])
}
