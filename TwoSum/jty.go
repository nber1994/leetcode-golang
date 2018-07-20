func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int, len(nums))
    res := make([]int, 2)
    for k, v := range nums {
        pom := target - v
        if index, exist := hashMap[pom]; exist{
            res[1] = k
            res[0] = index
            return res
        }
        hashMap[v] = k
    }
        return res
}

