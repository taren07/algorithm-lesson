package main

import (
	// "algorithm/src/recursive"
	// "algorithm/src/sort"
	"algorithm/src/euclidean"
	"fmt"
)

func main() {
	var A, B int64
	fmt.Scan(&A, &B)
	fmt.Println(euclidean.GCD(A, B))
}
