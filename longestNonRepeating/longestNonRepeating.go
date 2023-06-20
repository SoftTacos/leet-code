package longestnonrepeating

func lengthOfLongestSubstring(s string) (longestRun int) {
	var (
		exists        bool
		currentRun    int
		currentRunSet map[byte]struct{}
	)
	for i := 0; i < len(s)-longestRun; i++ { // can optimize this i < len(s) - longestRun
		// goto each char in string
		// start run off of that character
		currentRun = 0
		currentRunSet = map[byte]struct{}{}
		for j := i; j < len(s); j++ {
			// goto each character
			// if unique, add to currentRunSet and increment currentRun
			// if already exists in currentRunSet, check if currentRun > longestRun, set longestRun if >
			if _, exists = currentRunSet[s[j]]; exists {
				break
			}
			// fmt.Print(s[j : j+1])
			currentRunSet[s[j]] = struct{}{}
			currentRun++
			if currentRun > longestRun {
				longestRun = currentRun
			}
		}
		// fmt.Println(longestRun)
	}
	return
}
