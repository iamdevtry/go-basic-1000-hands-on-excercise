package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)

	if argCount == 1 {
		fmt.Println("Give me arguments.")
	} else if argCount == 2 {
		fmt.Println(`There is one: "hello"`)
	} else if argCount == 3 {
		fmt.Println(`There are two: "hi there`)
	} else {
		fmt.Printf("There are %d arguments\n", argCount-1)
	}
}
