package longest

func longestPalindrome(s string) (pal string) {
	// subs := make([]string, len(s)*len(s))
	// subs := make(map[string]struct{}, len(s)*len(s))
	var sub string
	for i := range s {
		for size := 1; size < len(s)-i+1; size++ {
			sub = s[i : i+size]
			// if sub is smaller than current largest palindrome, skip
			if len(sub) <= len(pal) {
				continue
			}

			// if we already tried that substring, skip
			if isPalindrome(sub) {
				// subs[s[i:i+size]] = struct{}{}
				if len(pal) < len(sub) {
					pal = sub
				}
			}
		}
	}

	return
}

func isPalindrome(subStr string) bool {
	if len(subStr) == 1 {
		return true
	}

	middle := len(subStr) / 2
	for i := 0; i < middle; i++ {
		if subStr[i] != subStr[len(subStr)-i-1] {
			return false
		}
	}
	return true
}
