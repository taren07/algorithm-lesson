// SelectionSort
function selectionSort(arr: number[]): void {
	const N: number = arr.length;
	for (let i = 0; i < N - 1; i++) {
		let minPosition: number = i;
		let minVal: number = arr[i];
		for (let j = i + 1; j < N; j++) {
			if (arr[j] < minVal) {
				minPosition = j;
				minVal = arr[j];
			}
		}
		[arr[i], arr[minPosition]] = [arr[minPosition], arr[i]];
	}
}
