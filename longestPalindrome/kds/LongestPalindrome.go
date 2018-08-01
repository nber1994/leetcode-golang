package kds

/**
 定位中心节点，向边界拓展
 每个节点中心节点有两个：mid和mid+1（因为字符串长度奇偶的关系）
 时间复杂度：O(n2)
 空间复杂度：O(1)
 */
func LongestPalindrome(s string) string {
	start, end := 0, 0
	length := len(s)
	if 0 == length {
		return s
	}
	for mid := 0; mid < length; mid++ {
		l1 := getMaxRound(s, length, mid, mid)
		l2 := getMaxRound(s, length, mid, mid+1)
		maxLen := end - start + 1
		if l1 > maxLen {
			start = mid-(l1-1)/2
			end = mid+(l1-1)/2
		}
		if l2 > maxLen {
			start = mid-l2/2+1
			end = mid+l2/2
		}
	}

	return string([]byte(s)[start:end+1])
}

func getMaxRound(s string, length int, left int, right int) int {
	for left >= 0 && right < length && s[left] == s[right] {
		right++
		left--
	}
	return right - left - 1
}
