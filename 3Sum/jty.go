func threeSum(nums []int) [][]int {
    lens := len(nums)
    res := [][]int{}
    sort.Ints(nums)
    for left := 0; left < lens - 2; left++ {
        if left > 0 && nums[left] == nums[left-1] {
            continue
        }
        mid := left + 1
        right := lens - 1
        for mid < right {
            if mid > left + 1 && nums[mid] == nums[mid - 1] {
                mid++
                continue
            }
            if nums[mid] + nums[left] + nums[right] == 0 {
                res = append(res, []int{nums[left], nums[mid], nums[right]})
                mid++
                right--
            } else if nums[mid] + nums[left] + nums[right] < 0 {
                mid++
            } else {
                right--
            }
        }
    }
    return res
}
