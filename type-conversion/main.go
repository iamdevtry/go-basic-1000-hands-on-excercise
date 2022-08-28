package main

import "fmt"

func main() {
	// 1
	// a, b := 10, 5.5
	// fmt.Println(float64(a) + b)

	// 2
	// a, b := 10, 5.5
	// a = int(b)
	// fmt.Println(float64(a) + b)

	// 3
	// fmt.Println(5.5)

	// 4
	// age := 2
	// fmt.Println(float64(7.5) + float64(age))

	//5
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))

}
