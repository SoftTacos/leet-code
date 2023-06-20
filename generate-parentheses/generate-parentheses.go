package generateparentheses

func generateParenthesis(n int) []string {
	return generate(n-1, n, "(")
}

func generate(open, close int, current string) (response []string) {
	if open == 0 && close == 0 {
		return []string{current}
	}
	if open > 0 {
		response = generate(open-1, close, current+"(")
	}
	if close > 0 && close > open {
		response = append(response, generate(open, close-1, current+")")...)
	}

	return
}
