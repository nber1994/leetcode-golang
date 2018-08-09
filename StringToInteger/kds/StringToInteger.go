package kds

func MyAtoi(str string) int {
	res, flag := 0, 1
	isStart := false
	for _, value := range str {
		if value >= '0' && value <= '9' {
			res = res*10 + int(value-'0')
			if res > 1<<31-1 {
				if 1 == flag {
					res = 1<<31 - 1
				} else {
					res = 1 << 31
				}
				break
			}
			isStart = true
		} else if !isStart {
			if '-' == value {
				flag = -1
			} else if '+' == value {
				flag = 1
			} else if ' ' == value {
				continue
			} else {
				break
			}
			isStart = true
		} else {
			break
		}
	}

	return res * flag
}
