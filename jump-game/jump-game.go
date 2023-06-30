package jumpgame

// [1,5,10,0,12,1]
func canJump(nums []int) (canJump bool) {
	canJumps := make([]bool, len(nums))
	canJumps[len(canJumps)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		// if nums[i] can just get us to the end of nums
		if i+nums[i] >= len(nums)-1 {
			canJumps[i] = true
			// otherwise look at the already calculated values in canJumps that nums[i] can get to
		} else if nums[i] > 0 {
			for j := 1; j <= nums[i]; j++ { // j is the offset from nums[i]
				if canJumps[i+j] {
					canJumps[i] = true
					break
				}
			}
		}
	}
	canJump = canJumps[0]
	return
}
