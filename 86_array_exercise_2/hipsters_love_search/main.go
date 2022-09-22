package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a book title")
		return
	}

	title := args[0]
	found := false
	for i, book := range books {
		if strings.Contains(strings.ToLower(book), strings.ToLower(title)) {
			fmt.Printf("#%d: %q\n", i+1, book)
			found = true
		}
	}

	if found == false {
		fmt.Printf("No book found with the title %q\n", title)
	}
}
