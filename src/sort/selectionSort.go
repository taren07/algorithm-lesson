package sort

// SelectionSort
func SelectionSort(arr []int) {
    N := len(arr)
    for i := 0; i < N-1; i++ {
        minPosition := i
        minVal := arr[i]
        for j := i + 1; j < N; j++ {
            if arr[j] < minVal {
                minPosition = j
                minVal = arr[j]
            }
        }
        arr[i], arr[minPosition] = arr[minPosition], arr[i]
    }
}
