package climbingstairs

func climbStairs(n int) (combinations int) {
	if n <= 3 {
		return n
	}
	// i = step, vals[i] = combinations
	vals := make([]int, n)
	vals[0] = 1
	vals[1] = 2
	vals[2] = 3
	for i := 3; i < n; i++ {
		vals[i] = vals[i-1] + vals[i-2] // this is the iterative breakdown from the recursive climbStars(n-1) climbstairs(n-2)
	}
	combinations = vals[n-1]
	return
}

func climbStairsRecursive(n int) (combinations int) {
	if n < 4 {
		return n
	}
	return climbStairs(n-1) + climbStairs((n - 2))
}
