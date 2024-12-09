package entry

import "fmt"

func FindMax(arr []int) int {
	max := arr[0]
	fmt.Println("max: ", max)

	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}