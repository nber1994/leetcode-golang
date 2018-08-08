package kds

func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	left1, right1, left2, right2 := 0, len1-1, 0, len2-1
	var index1, index2 []int
	if 0 == len1 {
		return getMedian(nums2, left2, right2)
	}
	if 0 == len2 {
		return getMedian(nums1, left1, right1)
	}
	for left1 < right1 && left2 < right2 {
		index1, index2 = getMedianIndex(left1, right1), getMedianIndex(left2, right2)
		if nums1[index1[0]] <= nums2[index2[1]] && nums2[index2[0]] <= nums1[index1[1]] {
			return float64(getMax(nums1[index1[0]], nums2[index2[0]])+getMin(nums1[index1[1]], nums2[index2[1]])) / 2
		}

		if nums1[index1[0]] > nums2[index2[1]] {
			abandon := getAbandonCount(left2, index2[1], right1, index1[0])
			left2 += abandon
			right1 -= abandon
		}
		if nums2[index2[0]] > nums1[index1[1]] {
			abandon := getAbandonCount(left1, index1[1], right2, index2[0])
			left1 += abandon
			right2 -= abandon
		}
	}

	if left1 == right1 && left2 == right2 {
		return float64(nums1[left1]+nums2[right2]) / 2
	} else {
		var left, right, challenger int
		var nums []int
		if left1 == right1 {
			left, right, challenger = left2, right2, nums1[left1]
			nums = nums2
		} else {
			left, right, challenger = left1, right1, nums2[left2]
			nums = nums1
		}
		index := getMedianIndex(left, right)
		if 0 == index[0] {
			return float64(getMid(nums[index[0]], nums[index[1]], challenger))
		} else {
			if nums[index[0]-1] > challenger {
				right--
				return getMedian(nums, left, right)
			} else if nums[index[0]+1] < challenger {
				left++
				return getMedian(nums, left, right)
			} else {
				if index[0] == index[1] {
					return float64(nums[index[0]]+challenger) / 2
				} else {
					return float64(getMid(nums[index[0]], nums[index[1]], challenger))
				}
			}
		}
	}
}

func getAbandonCount(oldLeft int, newLeft int, oldRight int, newRight int) int {
	abandon := 0
	if oldRight-newRight < newLeft-oldLeft {
		abandon = oldRight - newRight
	} else {
		abandon = newLeft - oldLeft
	}
	return abandon
}

func getMid(small int, big int, unknow int) int {
	if small >= unknow {
		return small
	} else if big <= unknow {
		return big
	} else {
		return unknow
	}
}

func getMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMedian(nums []int, left int, right int) float64 {
	index := getMedianIndex(left, right)
	return float64(nums[index[0]]+nums[index[1]]) / 2
}

func getMedianIndex(left int, right int) []int {
	sum := right + left
	return []int{sum / 2, (sum + 1) / 2}
}
