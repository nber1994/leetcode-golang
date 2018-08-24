package kds

/**
 * 暴力破解
 * 时间: O(n2)
 * 空间: O(1)
 */
func MaxArea(height []int) int {
	length := len(height)
	maxArea := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tmpHeight := height[i]
			if height[j] < height[i] {
				tmpHeight = height[j]
			}
			area := tmpHeight * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

/**
 * 双指针法
 * 时间: O(n1)
 * 空间：O(1)
 */
func MaxArea1(height []int) int {
	length, maxArea, minHeight, area := len(height), 0, 0, 0
	for i, j := 0, length-1; i < j; {
		if height[j] < height[i] {
			minHeight = height[j]
			j--
		} else {
			minHeight = height[i]
			i++
		}
		area = minHeight * (j - i + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
