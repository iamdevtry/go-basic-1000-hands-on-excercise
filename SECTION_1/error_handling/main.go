package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Success: Converted %q to %d.\n", age, n)
}
