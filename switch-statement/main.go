package main

import (
	"fmt"
	// "os"
	"time"
)

func main() {
	// city := os.Args[1]
	// switch city {
	// case "Paris":
	// 	fmt.Println("France")
	// 	fallthrough
	// case "Tokyo":
	// 	fmt.Println("Japan")
	// }
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
