package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age.")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if age < 0 {
		fmt.Println("Error: Age cannot be negative.")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	} else if age <= 17 && age >= 13 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("R-Rated")
	}
}
