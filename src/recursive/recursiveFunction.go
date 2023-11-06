package recursive

func Factorial(N int) int {
	if N == 1 {
		return 1 // Cases like this that should be distinguished are called "base cases."
	}
	return Factorial(N-1) * N
}