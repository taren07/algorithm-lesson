package sort

// SelectionSort
func SelectionSort(nums []int) []int {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if nums[j] < nums[minIndex] {
                minIndex = j
            }
        }

        if i != minIndex {
            nums[i], nums[minIndex] = nums[minIndex], nums[i]
        }
    }

    return nums
}

// Please write the main function below.

// import (
// 	"algorithm/src/sort"
// 	"fmt"
// )

// func main() {
// arr := []int{12, 4, 7, 2, 10, 8}
// sort.SelectionSort(arr)
// var N int
// fmt.Print("Enter a number: ")
// fmt.Scan(&N)
// result := recursive.Factorial(N)
// fmt.Println(result)
// }
