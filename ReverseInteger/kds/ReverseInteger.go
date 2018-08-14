package kds

func Reverse(x int) int {
	var integer, flag int
	flag = 1
	if x < 0 {
		flag = -1
		x = flag * x
	}
	for integer = 0; x/10 > 0; x = x / 10 {
		if 0 == integer && 0 == x%10 {
			continue
		}
		integer = integer*10 + x%10
	}
	integer = integer*10 + x%10
	if integer > 2 << 31 - 1 {
		integer = 0
	}
	return flag * integer
}

func Reverse1(x int32) int32 {
	var integer, flag int32
	flag = 1
	if x < 0 {
		flag = -1
		x = flag * x
	}
	for integer = 0; x / 10 > 0; x = x / 10 {
		if 0 == integer && 0 == x % 10 {
			continue
		}
		integer = integer * 10 + x % 10
	}
	integer = integer * 10 + x % 10
	if integer < 0 {
		integer = 0
	}
	return flag * integer
}