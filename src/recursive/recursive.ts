function factorial(N: number): number {
	if (N === 1) {
		return 1; // This kind of case distinction is called a "base case."
	}
	return factorial(N - 1) * N;
}
const userInput = prompt("Enter a number:");
if (userInput !== null) {
	const N: number = parseInt(userInput);
	console.log(factorial(N));
} else {
	console.log("No number was entered.");
}
