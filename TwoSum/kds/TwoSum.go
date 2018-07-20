package kds

import (
	"sort"
)

func TwoSum1(nums []int, target int) []int {
	count := len(nums)
	hashMap := make(map[int]int, count)
	res := make([]int, 2)
	for k, v := range nums {
		index, exist := hashMap[target - v]
		if !exist {
			hashMap[v] = k
			continue
		}
		res[0] = index
		res[1] = k
	}
	return res
}

func TwoSum2(nums []int, target int) []int {
	count := len(nums)
	tmpNums := make([]int, count)
	copy(tmpNums[:], nums)
	res := []int{}
	i := 0
	j := count - 1
	sort.Ints(nums)  // 从小到大，排序

	for i < j {
		sum := nums[i] + nums[j]
		if sum > target {
			j--
			continue
		} else if sum < target {
			i++
			continue
		} else {
			break
		}
	}
	if i < j {
		for k, v := range tmpNums {
			if nums[i] == v || nums[j] == v {
				res = append(res, k)
			}
		}
	}

	return res
}