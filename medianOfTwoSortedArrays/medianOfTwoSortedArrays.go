package medianOfTwoSortedArrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) (median float64) {
	numsTogetherStrong := append(nums1, nums2...)
	sort.Ints(numsTogetherStrong)
	if len(numsTogetherStrong) == 0 {
		return
	} else if len(numsTogetherStrong)%2 == 1 {
		//odd
		median = float64(numsTogetherStrong[len(numsTogetherStrong)/2])
	} else {
		//even
		median = float64(numsTogetherStrong[len(numsTogetherStrong)/2-1]+numsTogetherStrong[len(numsTogetherStrong)/2]) / 2
	}
	return
}
