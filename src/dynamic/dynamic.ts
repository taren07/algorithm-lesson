// Input
const [N, S] = readLine().split(" ").map(Number); // Read N and S
const A = readLine().split(" ").map(Number); // Read array A

// Initialize array
const dp: (boolean | null)[][] = Array.from({ length: N + 1 }, () =>
	Array(S + 1).fill(null)
);
dp[0][0] = true;
for (let i = 1; i <= S; i++) {
	dp[0][i] = false;
}

// Dynamic programming
for (let i = 1; i <= N; i++) {
	for (let j = 0; j <= S; j++) {
		if (j < A[i - 1]) {
			// When j < A[i-1], cannot choose card i
			dp[i][j] = dp[i - 1][j];
		} else {
			// When j >= A[i-1], there are two options: choose or not choose
			dp[i][j] = dp[i - 1][j] || dp[i - 1][j - A[i - 1]];
		}
	}
}

// Output the answer
if (dp[N][S]) {
	console.log("Yes");
} else {
	console.log("No");
}
