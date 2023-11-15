package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	// Initialize variables to store the closest sum and the minimum difference
	var (
		minDiff           = math.MaxInt32
		closestSum, sum   int
		left, right, diff int
	)
	// Iterate through the array
	for i := 0; i < len(nums)-2; i++ {
		// Two pointers approach for the remaining elements
		left, right = i+1, len(nums)-1

		for left < right {
			// Calculate the sum of the three integers
			sum = nums[i] + nums[left] + nums[right]

			// Calculate the absolute difference between the sum and the target
			diff = int(math.Abs(float64(sum - target)))

			// Update the closest sum and minimum difference if a better solution is found
			if diff < minDiff {
				minDiff = diff
				closestSum = sum
			}

			// Adjust pointers based on the comparison with the target
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}
