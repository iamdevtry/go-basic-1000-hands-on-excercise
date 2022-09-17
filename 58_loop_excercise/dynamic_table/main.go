package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	//header
	fmt.Printf("%-5s", "x")
	for i := 0; i <= num; i++ {
		fmt.Printf("%-5d", i)
	}
	fmt.Println()

	//column
	for i := 0; i <= num; i++ {
		fmt.Printf("%-5d", i)
		for j := 0; j <= num; j++ {
			fmt.Printf("%-5d", i*j)
		}
		fmt.Println()
	}

}
