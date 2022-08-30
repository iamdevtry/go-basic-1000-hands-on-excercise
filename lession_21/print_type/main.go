package main

import "fmt"

func main() {
	fmt.Printf("value: %d\ntype: %[1]T\n", 3)
	fmt.Printf("value: %.2f\ntype: %[1]T\n", 3.14)
	fmt.Printf("value: %s\ntype: %[1]T\n", "Hello, World!")
	fmt.Printf("value: %t\ntype: %[1]T\n", true)
}
