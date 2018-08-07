package jty

//窗口
func lengthOfLongestSubstring(s string) int {
	strArr := []byte(s)
	byteMap := make(map[byte]int, 0)
	lens := len(strArr)
	maxLen := 0
	left := 0
	right := 0
	for right < lens {
		if _, exist := byteMap[strArr[right]]; exist {
			if byteMap[strArr[right]] >= left {
				left = byteMap[strArr[right]] + 1
			}
		}
		byteMap[strArr[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}

//时间复杂度O(n)
//空间复杂度O(n)
