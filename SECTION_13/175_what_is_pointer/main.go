package main

import "fmt"

func main() {
	var counter byte = 100
	P := &counter
	V := *P
	*P = 25
	fmt.Println("counter: ", counter)
	fmt.Println("address of counter: ", &counter)

	fmt.Println("address of P: ", &P)
	fmt.Println("value of P: ", *P)

	fmt.Println("address of V: ", &V)
	fmt.Println("value of V: ", V)
}
