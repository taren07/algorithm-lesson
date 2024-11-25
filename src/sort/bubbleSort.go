package sort

func BubbleSort(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i ++ {
		for j := 0; j < n-i-1; j ++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// import (
// 	"algorithm/src/sort"
// 	"fmt"
// )

// nums := []int{15, 38, 2, 21}
// 	sort.BubbleSort(nums)
// 	fmt.Println(nums)