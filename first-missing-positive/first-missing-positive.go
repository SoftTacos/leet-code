package firstmissingpositive

func firstMissingPositive(nums []int) int {
	var (
		vals   = map[int]bool{}
		maxVal int
	)
	// goto each array
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
		if n > 0 {
			vals[n] = true
		}
	}

	// find element i where false, add 1
	// for i := range vals {
	for i := 1; i < maxVal; i++ {
		if !vals[i] {
			return i
		}
	}

	return maxVal + 1
}
