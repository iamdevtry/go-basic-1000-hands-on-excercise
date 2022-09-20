package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		println("Give me your name!")
		return
	}

	name := os.Args[1]
	//declare elipsis string array
	moody := [...]string{
		"Sad",
		"Happy",
		"Excited",
		"Angry",
		"Surprised",
		"Confused",
		"Disgusted",
		"Scared",
	}

	//random a number
	rand := rand.Intn(len(moody))
	fmt.Printf("%s is %s.\n", name, moody[rand])
}
