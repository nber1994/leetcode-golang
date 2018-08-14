package kds

/**
 https://cloud.tencent.com/developer/article/1173458
 */
func Convert(s string, numRows int) string {
	length := len(s)
	batchCount := 2 * (numRows - 1)
	if 1 >= numRows || length < numRows {
		return s
	}
	batchTotal := length / batchCount
	lastBatchCount := length % batchCount
	i := 0
	if 0 != lastBatchCount {
		batchTotal++
	}
	newStr := make([]byte, length)

	for ; i < batchTotal; i++ {
		newStr[i] = s[i*2*(numRows-1)]
	}
	for row := 2; row < numRows; row++ {
		index := row - 1
		delta1, delta2 := batchCount-2*row+2, 2*row-2
		for j := 1; index < length; i++ {
			newStr[i] = s[index]
			if 1 == j%2 {
				index += delta1
			} else {
				index += delta2
			}
			j++
		}
	}
	index := numRows-1
	for ; i < length && index < length; i++ {
		newStr[i] = s[index]
		index += batchCount
	}

	return string(newStr)
}
