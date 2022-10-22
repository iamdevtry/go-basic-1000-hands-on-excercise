package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	num, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	//header
	op := os.Args[1]
	fmt.Printf("%-5s", op)
	for i := 0; i <= num; i++ {
		fmt.Printf("%-5d", i)
	}
	fmt.Println()

	//column
	for i := 0; i <= num; i++ {
		fmt.Printf("%-5d", i)
		for j := 0; j <= num; j++ {
			switch op {
			case "+":
				fmt.Printf("%-5d", i+j)
			case "-":
				fmt.Printf("%-5d", i-j)
			case "*":
				fmt.Printf("%-5d", i*j)
			case "/":
				fmt.Printf("%-5d", i/j)
			default:
				fmt.Println("Please enter a valid operator")
				return
			}
		}
		fmt.Println()
	}

}
