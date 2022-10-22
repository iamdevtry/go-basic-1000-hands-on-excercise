package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [command] [string]")
		return
	}

	command, str := os.Args[1], os.Args[2]
	switch command {
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "title":
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command %q.\n", command)
	}
}
