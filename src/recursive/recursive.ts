function factorial(N: number): number {
	if (N === 1) {
		return 1; // This kind of case distinction is called a "base case."
	}
	return factorial(N - 1) * N;
}

const N: number = parseInt(prompt("Enter a number:"));
console.log(factorial(N));
