package sort

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
}

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{15, 38, 2, 21}
// 	sort.InsertionSort(arr)
// 	fmt.Println(arr)
// }
