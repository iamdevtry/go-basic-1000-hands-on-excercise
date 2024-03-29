package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n   int
		err error
	)

	if len(os.Args) != 2 {
		fmt.Println("Please input a number")
		return
	} else if n, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}
