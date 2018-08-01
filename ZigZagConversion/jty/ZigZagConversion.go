package jty

func convert(s string, numRows int) string {
	sArr := []byte(s)
	if len(sArr) < 2 || numRows == 1 {
		return s
	}
	resArr := []byte{}
	step := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(sArr); j = j + step {
			resArr = append(resArr, sArr[j+i])
			if i != 0 && i != numRows-1 && j+step-i < len(sArr) {
				resArr = append(resArr, sArr[j+step-i])
			}
		}
	}
	return string(resArr)
}

//时间复杂度O(n)
//空间复杂度O(n)
