package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usageMsg         = "Usage [min] [max]"
	minBiggerThanMax = "Min cannot bigger than max"
	numberInvalidMsg = "Number is invalid"
)

func main() {
	if len(os.Args) != 3 {
		println(usageMsg)
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	max, err := strconv.Atoi(os.Args[2])

	if err != nil {
		println(numberInvalidMsg)
		return
	}

	if min > max {
		println(minBiggerThanMax)
		return
	}

	sum := 0
	for i := min; i <= max; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum from", min, "to", max, "is", sum)
}
