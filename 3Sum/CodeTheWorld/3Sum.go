package CodeTheWorld

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	sum := 0
	var res [][]int

	for i := 0; i < length-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, length-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else if sum < 0 {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}
		}
	}

	return res
}
