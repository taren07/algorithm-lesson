package euclidean

func GCD(A, B int64) int64 {
	for A >= 1 && B >= 1 {
		if A < B {
			B = B % A
		} else {
			A = A % B
		}
	}
	if A >= 1 {
		return A
	}
	return B
}


// Put the following in the main file and run it

// import (
// 	"algorithm/src/euclidean"
// 	"fmt"
// )

// var A, B int64
// 	fmt.Scan(&A, &B)
// 	fmt.Println(euclidean.GCD(A, B))