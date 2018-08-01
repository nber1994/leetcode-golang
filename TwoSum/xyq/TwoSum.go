package main

import "fmt"

func main() {
    nums := []int{3, 2, 4}
    target := 6
    res := twoSum2(nums, target);
    for _, num := range res {
        fmt.Println(':', num)
    }
}

func twoSum1(nums []int, target int) []int {
    for i, num := range nums {
        for j := i + 1; j < len(nums); j++ {
            if num + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return []int{}
}

func twoSum2(nums []int, target int) []int {
    v_k := make(map[int]int)
    for i, num := range nums {
        v_k[num] = i
    }
    for j, num := range nums {
        fmt.Println(num)
        res := target - num
        k, ok := v_k[res]
        if (ok && k != j) {
            return []int{j, k}
        }
    }

    return []int{}
}
