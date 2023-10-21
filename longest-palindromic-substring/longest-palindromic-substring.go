package longest

func longestPalindrome(s string) (pal string) {
	subs := make([]string, 0, len(s)*len(s))
	// fmt.Println(s)
	for i := range s {
		for size := 1; i+size < len(s)+1; size++ {
			subs = append(subs, s[i:i+size])
		}
	}

	for s := range subs {
		if len(subs[s]) <= len(pal) {
			continue
		} else if isPalindrome(subs[s]) {
			pal = subs[s]
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
