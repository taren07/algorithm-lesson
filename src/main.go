package main

import (
	// "algorithm/src/recursive"
	// "algorithm/src/sort"
	// "algorithm/src/euclidean"
	"algorithm/src/sort"
	"fmt"
)

func main() {
	arr := []int{4, 7, 2, 5, 8, 1, 3, 6, 9}
	sort.MergeSort(arr)
	fmt.Println(arr)
}
