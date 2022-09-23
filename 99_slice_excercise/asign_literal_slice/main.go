package main

import "fmt"

func main() {
	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alives    []bool
	)

	names = []string{"serpil", "ebru", "lina"}
	distances = []int{100, 200, 300, 400, 500}
	data = []byte{'I', 'N', 'A', 'N', 'C'}
	ratios = []float64{3.14, 6.28}
	alives = []bool{true, false, false, true}

	fmt.Printf("%T %d %v\n", names, len(names), names == nil)
	fmt.Printf("%T %d %v\n", distances, len(distances), distances == nil)
	fmt.Printf("%T %d %v\n", data, len(data), data == nil)
	fmt.Printf("%T %d %v\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("%T %d %v\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances and the data slices are the same.")
	}
}
