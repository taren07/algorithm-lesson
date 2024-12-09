function findMax(arr: number[]): number {
	let max = arr[0];
	for (const num of arr) {
		if (num > max) {
			max = num;
		}
	}
	return max;
}

const nums = [1, 5, 3, 9, 2];
console.log(findMax(nums));
