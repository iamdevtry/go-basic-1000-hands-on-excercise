package main

import (
	"fmt"
)

func main() {
	// Please uncomment the code below

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	// for _, word := range words {
	// 	for _, w := range []byte(word) {
	// 		fmt.Print(w, " ")
	// 	}
	// 	fmt.Println()

	// 	bwords = append(bwords, []byte(word))
	// }
	// fmt.Println()
	// for _, bword := range bwords {
	// 	for _, w := range bword {
	// 		fmt.Printf("%s", string(w))
	// 	}
	// 	fmt.Println()
	// }

	//SOLUTION
	for _, w := range words {
		bw := []byte(w)

		fmt.Println(bw)

		bwords = append(bwords, bw)
	}

	for _, w := range bwords {
		fmt.Println(string(w))
	}
}
