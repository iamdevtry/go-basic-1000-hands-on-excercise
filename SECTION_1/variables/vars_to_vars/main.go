package main

import "fmt"

func main() {
	color := "green"
	darkColor := "dark"
	color = darkColor + " " + color
	fmt.Println(color)
}
