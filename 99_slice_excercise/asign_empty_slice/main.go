package main

import "fmt"

func main() {
	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var alives []bool

	names = []string{}
	distances = []int{}
	data = []byte{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("%T %d %v\n", names, len(names), names == nil)
	fmt.Printf("%T %d %v\n", distances, len(distances), distances == nil)
	fmt.Printf("%T %d %v\n", data, len(data), data == nil)
	fmt.Printf("%T %d %v\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("%T %d %v\n", alives, len(alives), alives == nil)
}
