package kds

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	newNumber := 0
	for ; tmp > 0; tmp /= 10 {
		newNumber = newNumber*10 + tmp%10
	}
	return newNumber == x
}
