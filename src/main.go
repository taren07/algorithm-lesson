package main

import (
	"algorithm/src/recursive"
	"algorithm/src/sort"
	"fmt"
)

func main() {
	arr := []int{12, 4, 7, 2, 10, 8}

	// SelectionSort
	sort.SelectionSort(arr)

	var N int
	fmt.Print("Enter a number: ")
	fmt.Scan(&N)
	result := recursive.Factorial(N)
	fmt.Println(result)
}
