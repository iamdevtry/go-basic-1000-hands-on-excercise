package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", arg)
		return
	}
	meters := feet * 0.3084

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
