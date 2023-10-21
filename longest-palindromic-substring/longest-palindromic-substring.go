package longest

func longestPalindrome(s string) (pal string) {
	// subs := make(string, len(s)*len(s))
	subs := make(map[string]struct{}, len(s)*len(s))
	// fmt.Println(s)
	var exists bool
	var sub string
	for i := range s {
		for size := 1; size < len(s)-i+1; size++ {
			sub = s[i : i+size]
			// if sub is smaller than current largest palindrome, skip
			// fmt.Println(sub, i, size)
			if len(sub) <= len(pal) {
				continue
			}

			// if we already tried that substring, skip
			if _, exists = subs[sub]; exists {
				continue
			} else if isPalindrome(sub) {
				subs[s[i:i+size]] = struct{}{}
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
	offset := 1
	if len(subStr)%2 == 0 {
		offset = 0
	}
	left := subStr[:len(subStr)/2]
	right := subStr[len(subStr)/2+offset:]
	// fmt.Printf("%s:%s%s\n", subStr, left, right)
	for i := range left {
		if left[i] != right[len(right)-i-1] {
			return false
		}
	}
	return true
}

/*

func longestPalindrome(s string) (pal string) {
	if s == "" {
		return s
	}
	pal = s[0:1]

	lengths := make([]int, len(s))
	var start, end int
	for i := range s {
		// beginning or end of slice
		if i == 0 || i == len(s)-1 {
			lengths[i] = 1
			continue
		}

		for radius := 0; radius < len(s)-i-1; radius++ {
			if s[i] == s[i+1] {
				start = i
				end = i + 1
			} else if s[i-1] == s[i+1] {
				start = i - 1
				end = i + 1
			}
			fmt.Println(s, s[start:end+1], start, end)
			if isPalindrome(s[start : end+1]) {
				fmt.Println("is")
			}
		}
	}

	return pal
}

func isPalindrome(subStr string) bool {
	if len(subStr) == 1 {
		return true
	}
	offset := 1
	if len(subStr)%2 == 0 {
		offset = 0
	}
	left := subStr[:len(subStr)/2]
	right := subStr[len(subStr)/2+offset:]
	fmt.Printf("%s:%s%s\n", subStr, left, right)
	for i := range left {
		if left[i] != right[len(right)-i-1] {
			return false
		}
	}
	return true
}


*/
