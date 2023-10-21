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
