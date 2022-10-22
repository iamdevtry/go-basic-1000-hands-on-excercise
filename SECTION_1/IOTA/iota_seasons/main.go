package main

import "fmt"

func main() {
	const (
		winter = (iota + 1) * 3
		spring
		summer
		fall
	)

	fmt.Println(winter, spring, summer, fall)
}
