package main

import "fmt"

func main() {
	week := [...]string{"Monday", "Tuesday"}
	wend := [2]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
