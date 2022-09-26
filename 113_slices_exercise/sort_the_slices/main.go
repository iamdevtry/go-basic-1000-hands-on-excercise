package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	mid := len(items) / 2
	itemsCopy := items[mid-1 : mid+1]
	sort.Strings(itemsCopy)
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
