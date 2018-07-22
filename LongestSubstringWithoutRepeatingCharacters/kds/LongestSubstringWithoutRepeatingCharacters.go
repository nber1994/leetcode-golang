package kds

func LengthOfLongestSubstring1(s string) int {
	strArr := []byte(s)
	strLen := len(strArr)
	maxLen := 0

	for left := range strArr {
		right := 0
		charArr := [128]int{}
		for index := left; index < strLen; index++ {
			charIndex := strArr[index]
			if charArr[charIndex] > 0 {
				break
			}
			right = index + 1
			charArr[charIndex]++
		}
		currentLen := right - left
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func LengthOfLongestSubstring2(s string) int {
	strArr := []byte(s)
	length := len(strArr)
	left, right, maxLen := 0, 0, 0
	charArr := [128]int{}

	for right < length {
		charIndex := s[right]
		if charArr[charIndex] > left {
			currentLen := right - left
			if currentLen > maxLen {
				maxLen = currentLen
			}
			left = charArr[charIndex]
		}
		right++
		charArr[charIndex] = right
	}
	lastLen := length - left
	if lastLen > maxLen {
		maxLen = lastLen
	}

	return maxLen
}
