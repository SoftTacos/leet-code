package longest

func longestPalindrome(s string) (pal string) {
	var current string
	for offset := range s {
		current = longestFromOffset(s, offset)
		if len(current) > len(pal) {
			pal = current
		}
	}
	return
}

func longestFromOffset(s string, offset int) (longest string) {
	if offset == 0 || offset > len(s)-1 {
		return s[offset : offset+1]
	}
	left := s[:offset] // if s == asdf then left == as

	// odd
	longest = calcLongest(left, s[offset:], "")

	// even
	if offset+1 == len(s) {
		return
	}
	evenLongest := calcLongest(left, s[offset+1:], s[offset:offset+1])
	if len(evenLongest) > len(longest) {
		longest = evenLongest
	}

	return
}

func calcLongest(left, right, mid string) string {
	var r int
	var l = len(left)
	for i := 0; i < len(left) && i < len(right); i++ {
		if left[len(left)-i-1] == right[i] {
			l = len(left) - i - 1
			r = i + 1
		} else {
			break
		}
	}
	return left[l:] + mid + right[:r]
}

/*

func longestFromOffset(s string, offset int) (longest string) {
	if offset >= len(s)-1 {
		return s[offset : offset+1]
	}

	longest = s[offset : offset+1]
	if s[offset] == s[offset+1] {
		longest = s[offset : offset+2]
	}
	var mod int
	for radius := 1; offset+radius < len(s) && offset-radius >= 0; radius++ {
		if len(longest)%2 == 0 {
			mod = 1
		} else {
			mod = 0
		}
		if offset+radius+mod < len(s) && s[offset-radius] == s[offset+radius+mod] {
			longest = s[offset-radius : offset+radius+mod+1]
		} else {
			return
		}
	}

	return
}

*/
