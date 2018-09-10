package codetheworld

func RomanToInt(s string) int {
	result := 0
	for index, value := range s {
		switch value {
		case 'I':
			result += 1
		case 'V':
			result += 5
			if index > 0 && 'I' == s[index-1] {
				result -= 2
			}
		case 'X':
			result += 10
			if index > 0 && 'I' == s[index-1] {
				result -= 2
			}
		case 'L':
			result += 50
			if index > 0 && 'X' == s[index-1] {
				result -= 20
			}
		case 'C':
			result += 100
			if index > 0 && 'X' == s[index-1] {
				result -= 20
			}
		case 'D':
			result += 500
			if index > 0 && 'C' == s[index-1] {
				result -= 200
			}
		case 'M':
			result += 1000
			if index > 0 && 'C' == s[index-1] {
				result -= 200
			}
		}
	}
	return result
}
