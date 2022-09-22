package main

import (
	"fmt"
	"strings"
)

func main() {
	people := [5][3]string{
		{"Albert", "Einstein", "Time"},
		{"Isaac", "Newton", "Apple"},
		{"Marie", "Curie", "Radium"},
		{"Charles", "Darwin", "Fittest"},
		{"Stephen", "Hawking", "Black Hole"},
	}

	//header
	fmt.Printf("%-15v %-15v %-15v\n", "First Name", "Last Name", "Nickname")
	fmt.Printf(strings.Repeat("=", 40) + "\n")
	for _, person := range people {
		fmt.Printf("%-15v %-15v %-15v\n", person[0], person[1], person[2])
	}

}
