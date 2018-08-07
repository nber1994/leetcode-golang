package jty

import (
	"bytes"
	"strings"
)

//manacher算法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	byteArr := []byte(s)
	byteArrArr := bytes.Split(byteArr, []byte{})
	sArr := bytes.Join(byteArrArr, []byte{'#'})
	sArr = append(sArr, 35)
	sArr = append([]byte{'#'}, sArr...)
	maxRight, mid, maxP, miden := 0, 0, 0, 0
	p := map[int]int{}
	for i := 1; i < len(sArr)-1; i++ {
		if i < maxRight {
			if p[2*mid-i] < maxRight-i {
				p[i] = p[2*mid-i]
			} else {
				p[i] = maxRight - i
			}
		} else {
			p[i] = 1
		}
		for i+p[i] <= len(sArr)-1 && i-p[i] >= 0 && sArr[i+p[i]] == sArr[i-p[i]] {
			p[i]++
		}
		if i+p[i] > maxRight {
			maxRight = i + p[i]
			mid = i
		}
		if maxP < p[i] {
			maxP = p[i]
			miden = i
		}
	}
	return strings.Replace(string(sArr[miden-maxP+1:miden+maxP]), "#", "", -1)
}

//时间复杂度O(n)
//空间复杂度O(n)
