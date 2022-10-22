package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	richter, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch {
	case richter < 2.0:
		fmt.Println("Micro")
	case richter >= 2.0 && richter < 3.0:
		fmt.Println("Very minor")
	case richter >= 3.0 && richter < 4.0:
		fmt.Println("Minor")
	case richter >= 4.0 && richter < 5.0:
		fmt.Println("Light")
	case richter >= 5.0 && richter < 6.0:
		fmt.Println("Moderate")
	case richter >= 6.0 && richter < 7.0:
		fmt.Println("Strong")
	case richter >= 7.0 && richter < 8.0:
		fmt.Println("Major")
	case richter >= 8.0 && richter < 10.0:
		fmt.Println("Great")
	case richter >= 10.0:
		fmt.Println("Meteoric")
	}
}
