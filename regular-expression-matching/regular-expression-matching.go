package regularexpressionmatching

import "fmt"

type Token struct {
	T    int
	Val  rune
	Next *Token
}

// func (t *Token) Matches(c byte) bool {
// 	if t.T == 0 {
// 		return c == t.Val
// 	} else if t.T == 1 {
// 		return true
// 	} else if t.T == 2 {
// 		return
// 	}
// }

func isMatch(str, pattern string) bool {
	pN := tokenize(pattern)
	sN := tokenize(str)

	for {
		if pN == nil || sN == nil {
			break
		}
		p := pN.Val
		s := sN.Val
		fmt.Println(string(s), string(p))
		switch pN.T {
		case 0, 1:
			if !simpleCharMatch(s, p) {
				return false
			}
			pN = pN.Next
			sN = sN.Next
		case 2:
			if simpleCharMatch(s, p) {
				sN = sN.Next
				// check next string character to see if it matches
			} else if pN.Next != nil && simpleCharMatch(s, pN.Next.Val) {
				pN = pN.Next
				if pN.Next != nil {
					pN = pN.Next
				}
				sN = sN.Next
			} else {
				return false
			}
		}
	}
	if pN == nil && sN == nil {
		return true
	}

	return false
}
func simpleCharMatch(c1, c2 rune) bool {
	if c1 == c2 || c2 == '.' {
		return true
	}
	return false
}

func tokenize(pattern string) (root *Token) {
	root = &Token{}
	current := root
	for _, p := range pattern {
		if byte(p) >= 97 && byte(p) <= 122 {
			// fmt.Println("REGULAR", string(p))
			// regular
			newToken := &Token{
				T:   0,
				Val: p,
			}
			current.Next = newToken
			current = newToken
		} else if p == '.' {
			// wildcard
			newToken := &Token{
				T:   1,
				Val: '.',
			}
			current.Next = newToken
			current = newToken
		} else if p == '*' {
			current.T = 2
		}
	}
	// for n := root; n != nil; n = n.Next {
	// 	fmt.Print(string(n.Val))
	// }
	// fmt.Println("")
	root = root.Next // TODO, unlazify this
	return
}

/*
func isMatch(str string, pattern string) bool {
	t := tokenize(pattern)
	for n := t; n != nil; n = n.Next {
		fmt.Print(string(n.Val))
	}
	fmt.Println("")
	for _, s := range str {
		if t == nil {
			fmt.Println("pattern empty")
			return false
		}
		fmt.Println(string(s), ":", string(t.Val))

		switch t.T {
		case 0, 1:
			if !simpleCharMatch(s, t.Val) {
				return false
			}
			t = t.Next
			// if i != len(str)-1 {
			// }
		case 2:
			if simpleCharMatch(s, t.Val) {
				continue
			} else if t.Next != nil && simpleCharMatch(s, t.Next.Val) {
				t = t.Next
				if t.Next != nil {
					t = t.Next
				}
				continue
			} else {
				return false
			}
		}
	}
	if t.Next != nil {
		return false
	}

	return true
}
*/

//@@@@@@@@@@@@@@@@@@@@@@@@@
// func isMatch(str string, pattern string) bool {
// 	var (
// 		// runChar byte
// 		s, p   byte
// 		pIndex int
// 	)

// 	for i := 0; i < len(str); i++ {
// 		s = str[i]
// 		if pIndex > len(pattern)-1 {
// 			return false
// 		}
// 		p = pattern[pIndex] //todo
// 		fmt.Println(string(s), string(p), pIndex)
// 		// if it's a regular match, advance s & p
// 		// if it's a . wildcard match, advance s & p
// 		if p != '*' {
// 			pIndex++
// 			if simpleCharMatch(s, p) {
// 				continue
// 			} else {
// 				return false
// 			}
// 		}

// 		// otherwise there be some * stuff
// 		// if it's a * match, advance s but not p
// 		// if s matches the repeating p character, continue
// 		if simpleCharMatch(s, pattern[pIndex-1]) { //s == pattern[pIndex-1] || pattern[pIndex-1] == '.' {
// 			continue
// 			// otherwise check if s matches the next p character, which would mean the "run" has ended
// 		} else if simpleCharMatch(s, pattern[pIndex+1]) {
// 			fmt.Println("SKIP", string(s), string(pattern[pIndex]), string(pattern[pIndex+1]))
// 			pIndex += 2
// 			continue
// 			// otherwise it doesn't match
// 		} else {
// 			return false
// 		}
// 	}

// 	// check if there are any remaining characters in the pattern
// 	// TODO: handle * endings
// 	if pIndex < len(pattern) && p != '*' {
// 		return false
// 	}
// 	return true
// }

// // pattern char must go second
// func simpleCharMatch(c1, c2 byte) bool {
// 	if c1 == c2 || c2 == '.' {
// 		return true
// 	}
// 	return false
// }
