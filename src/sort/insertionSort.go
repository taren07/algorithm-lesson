package sort

import (
	"fmt"
)

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j :=  i - 1
		fmt.Println(j)
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			fmt.Println(nums)
			j--
		}
		nums[j+1] = key
	}
	return nums
}

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{15, 38, 2, 21}
// 	sort.InsertionSort(arr)
// 	fmt.Println(arr)
// }
