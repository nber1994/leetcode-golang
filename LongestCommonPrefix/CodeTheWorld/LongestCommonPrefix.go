package CodeTheWorld

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	arrayLen := len(strs)
	prefixLen := len(strs[0])
	j := 0
	for i := 1; i < arrayLen; i++ {
		strLen := len(strs[i])
		for j = 0; j < strLen && j < prefixLen; j++ {
			fmt.Println(j)
			if strs[i-1][j] != strs[i][j] {
				break
			}
		}
		prefixLen = j
	}
	if 0 == prefixLen {
		return ""
	}
	return strs[0][0:prefixLen]
}
