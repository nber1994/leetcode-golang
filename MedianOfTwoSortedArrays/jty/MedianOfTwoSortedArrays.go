package jty

//在数组中虚拟加入#
//1，4，5，7 => 1,#,4,#,5,#,7,#
//这样以来，数组固定为奇数，避免了数组是偶数时，没有确切中位数index的尴尬
//此时，分割位置c = (2len + 1)/2 + 1, 左边的index为l = (c - 1) / 2, 右边的index为r = c/2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if n == 0 {
		return float64(nums2[m/2]+nums2[(m-1)/2]) / 2
	}
	if m == 0 {
		return float64(nums1[n/2]+nums1[(n-1)/2]) / 2
	}
	if n > m {
		//保证nums1是最小的
		return findMedianSortedArrays(nums2, nums1)
	}
	//定义最小的int和最大的int
	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)
	minInt := -maxInt - 1
	l1, l2, r1, r2, l, r, c1, c2 := 0, 0, 0, 0, 0, 0, 0, 0
	//两个边界指针
	left := 0
	right := 2*n + 1 - 1
	//二分长度较短的数组
	for left <= right {
		c1 = (left + right) / 2
		c2 = (2*n+2*m)/2 - c1
		if c1 == 0 {
			l1 = minInt
		} else {
			l1 = nums1[(c1-1)/2]
		}
		if c1 == 2*n {
			r1 = maxInt
		} else {
			r1 = nums1[c1/2]
		}
		if c2 == 0 {
			l2 = minInt
		} else {
			l2 = nums2[(c2-1)/2]
		}
		if c2 == 2*m {
			r2 = maxInt
		} else {
			r2 = nums2[c2/2]
		}

		if l1 > r2 {
			right = c1 - 1
		} else if l2 > r1 {
			left = c1 + 1
		} else {
			break
		}

	}
	//获取中位数
	if l1 > l2 {
		l = l1
	} else {
		l = l2
	}
	if r1 < r2 {
		r = r1
	} else {
		r = r2
	}
	return float64(r+l) / 2
}

//时间复杂度O(n+m)
//空间复杂度O(1)
//一遍遍历求中位数
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	mid := (len1 + len2) / 2
	index, i, j, miden, pre := 0, 0, 0, 0, 0
	for index <= mid {
		pre = miden
		if i == len1 {
			miden = nums2[j]
			j++
		} else if j == len2 {
			miden = nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				miden = nums1[i]
				i++
			} else {
				miden = nums2[j]
				j++
			}
		}
		index++
	}
	if (len1+len2)%2 == 1 {
		return float64(miden)
	} else {
		return float64(miden+pre) / 2
	}
}

//时间复杂度O((m+n)/2)
//空间复杂度O((m+n)/2)
