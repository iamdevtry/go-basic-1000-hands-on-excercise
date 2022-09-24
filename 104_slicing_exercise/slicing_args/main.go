package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please provide a string - Usage: start end")
		return
	}
	start, err1 := strconv.Atoi(args[0])
	end, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number")
		return
	}

	if start < 0 || end > len(ships) {
		fmt.Println("Out of range")
		return
	}

	newSlices := ships[start:end]

	fmt.Println(newSlices)
}
