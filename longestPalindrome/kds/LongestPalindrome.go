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
			start = mid - (l1-1)/2
			end = mid + (l1-1)/2
		}
		if l2 > maxLen {
			start = mid - l2/2 + 1
			end = mid + l2/2
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

/**
 Manacher算法，使用动态规划的思想
 Manacher算法讲解
	https://blog.csdn.net/hk2291976/article/details/51107886
	https://segmentfault.com/a/1190000003914228
 */
func LongestPalindrome2(s string) string {
	maxLen, maxPos, pos := 0, 0, 0
	maxStart, maxEnd := 0, 0
	newStr := initStr(s)
	length := len(newStr)
	strLenArr := make([]int, length)

	for i := 0; i < length; i++ {
		if i < maxPos {
			j := 2*pos - i
			if strLenArr[j] > maxPos-i {
				strLenArr[i] = maxPos - i
			} else {
				strLenArr[i] = strLenArr[j]
				continue
			}
		} else {
			strLenArr[i] = 1
		}
		for strLenArr[i]+i < length && i-strLenArr[i] >= 0 && newStr[i+strLenArr[i]] == newStr[i-strLenArr[i]] {
			strLenArr[i]++
		}
		maxPos = i + strLenArr[i] - 1
		pos = i
		if maxLen < strLenArr[i] {
			maxLen = strLenArr[i]
			maxStart = i - strLenArr[i] + 1
			maxEnd = i + strLenArr[i] - 1
		}
	}
	return string([]byte(s)[maxStart/2:maxEnd/2])
}

func initStr(s string) string {
	length := len(s)
	newStr := make([]byte, 2*length+1)
	for i, v := range s {
		newStr[2*i] = '#'
		newStr[2*i+1] = byte(v)
	}
	newStr[length*2] = '#'

	return string(newStr[:])
}
