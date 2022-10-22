package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("not a valid year\n")
		return
	}

	if year%400 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else if year%100 == 0 {
		fmt.Printf("%d is not a leap year.\n", year)
	} else if year%4 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
