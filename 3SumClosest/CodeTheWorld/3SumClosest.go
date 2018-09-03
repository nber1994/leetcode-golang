package CodeTheWorld

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	length, closest, gap := len(nums), 0, 0
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			tmp := sum - target
			if tmp < 0 {
				tmp *= -1
				j++
			} else if tmp > 0 {
				k--
			} else {
				return target
			}
			if gap == 0 || tmp < gap {
				gap = tmp
				closest = sum
			}
		}
	}

	return closest
}
