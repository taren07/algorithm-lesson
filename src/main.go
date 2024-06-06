package main

import (
	// "algorithm/src/recursive"
	// "algorithm/src/sort"
	// "algorithm/src/euclidean"
	"algorithm/src/sort"
	"fmt"
)

func main() {
	arr := []int{15, 38, 2, 21}
	sort.BubbleSort(arr)
	fmt.Println(arr)
}
