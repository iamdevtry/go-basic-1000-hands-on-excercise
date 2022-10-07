package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me a word")
		return
	}

	query := args[0]

	// var dictionary map[string]string
	// dictionary := map[string]string{}
	dictionary := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dictionary["up"] = "yukarı"
	dictionary["down"] = "aşağı"
	dictionary["mistake"] = ""

	value, ok := dictionary[query]
	if !ok {
		fmt.Printf("%q not found\n", query)
		return
	}

	fmt.Printf("%q means %#v\n", query, value)

	fmt.Printf("# of Keys: %d\n", len(dictionary))
	// fmt.Printf("Zero value: %#v\n", dictionary)
}
