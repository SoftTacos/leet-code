package integertoroman

// after writing this out I realize brute force would be better

var orderMapping = [4]map[int]byte{
	{
		1: 'I',
		5: 'V',
	},
	{
		1: 'X',
		5: 'L',
	},
	{
		1: 'C',
		5: 'D',
	},
	{
		1: 'M',
	},
}

func intToRoman(num int) (roman string) {
	var (
		res   = [][]byte{}
		order = 0
		digit int
		next  []byte
	)

	for num > 0 {
		digit = num % 10
		num = num / 10
		next = digitToRoman(digit, order)
		order++
		res = append(res, next)
	}

	next = nil
	for i := len(res) - 1; i >= 0; i-- {
		next = append(next, res[i]...)
	}
	roman = string(next)
	return
}

func digitToRoman(digit int, order int) (next []byte) {
	if digit == 4 {
		next = append(next, orderMapping[order][1])
		digit = 5
	}
	if digit == 9 {
		next = append(next, orderMapping[order][1])
		digit = 1
		order++
	}

	if digit >= 5 {
		next = append(next, orderMapping[order][5])
		digit -= 5
	}

	for digit > 0 {
		digit--
		next = append(next, orderMapping[order][1])
	}
	return
}
