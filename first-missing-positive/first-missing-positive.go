package firstmissingpositive

import (
	"math"
)

func firstMissingPositive(nums []int) int {
	var l = len(nums)

	// need to use the array to indicate if the value corresponding to its index has been used
	// if there is a missing value, the max missing value has to be [1,l], where l is the length of the list
	// if there is not a missing value, then return l+1

	// first toss out values that can't possibly be useful, because they are outside the possible values for a missing value
	// example: [1,-2,1,22,3] -> [1,6,1,6,3]
	// for non-go people,i is the index in the list, and n is a copy of the value at nums[i]
	for i, n := range nums {
		if n <= 0 || n > l {
			nums[i] = l + 1
		}
	}

	// now since every value in the list is either obviously l+1, or [1,l](no negatives)
	// i can safely ignore values equal to l+1. Now I need to indicate if a value exists or not
	// i can use that values index to mark its existence by making the value at the index negative
	// for non-go people,i is the index in the list, and n is a copy of the value at nums[i]
	// example continued: [1,6,1,5,3] -> [-1,6,-1,6,-3]
	for _, n := range nums {
		n = int(math.Abs(float64(n)))
		// skip values that are outside of l+1, we don't want to mess with that index
		if n == l+1 {
			continue
		}

		// set the index that represents the value(n-1) to be the negative value of whatever is there already
		// we want to preserve the value because that value could still be [1,l]
		nums[n-1] = -1 * int(math.Abs(float64(nums[n-1])))
	}

	// now I just find the first non-negative value
	// if the array is just 1,2,3,4,5, then all values will be negative, so l will still be equal to len(nums)
	for i, n := range nums {
		if n > 0 {
			l = i
			break
		}
	}
	return l + 1
}

// this was my initial solution, it passed the tests but it isn't actually
//  O(1) spacial complexity because I use a map to mark the existence of values
/*
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
*/
