package main

import (
	// "algorithm/src/recursive"
	// "algorithm/src/sort"
	// "algorithm/src/euclidean"
	// "algorithm/src/sort"
	"algorithm/src/entry"
	"fmt"
)

func main() {
	arr := []int{4, 7, 2, 5, 8, 1, 3, 6, 9}
	entry.FindMax(arr)
	fmt.Println(entry.FindMax(arr))
}
