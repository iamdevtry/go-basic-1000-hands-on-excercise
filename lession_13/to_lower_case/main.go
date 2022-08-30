package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	sLowerCase := strings.ToLower(s)
	fmt.Println(sLowerCase)
}
