package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("[your name] [positive|negative]]")
		return
	}

	name, mood := args[0], args[1]

	moodies := [][3]string{
		{"happy", "awesome", "great"},
		{"angry", "sad", "bad "},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moodies[0]))

	var mi int
	if mood == "positive" {
		mi = 0
	}

	fmt.Printf("%s is feeling %s today!", name, moodies[mi][n])
}
