package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := os.Args[1]
	l := len(a)
	repeatStr := strings.Repeat("!", l)
	fmt.Println(repeatStr + a + repeatStr)
}
