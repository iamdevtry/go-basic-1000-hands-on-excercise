package main

import "fmt"

func main() {
	fmt.Println("Sum numbers 1-10")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
