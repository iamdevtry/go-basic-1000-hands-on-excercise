package main

import (
	"fmt"
	"os"
)

func main() {

	// 1
	count := len(os.Args) - 1
	fmt.Printf("There are %d names.\n", count)

	// 2
	fmt.Println(os.Args[0])

	// 3
	fmt.Println(os.Args[1])

	// 4
	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")
}
