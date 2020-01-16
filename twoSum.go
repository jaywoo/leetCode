package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSumMap(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		val := target - nums[i]
		for j := 0; j < numsLen; j++ {
			if nums[j] == val && j != i {
				return []int{i, j}
			}
		}
	}

	return make([]int, 2)
}

func twoSumMap(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for k, v := range nums {
		numsMap[v] = k
	}

	for k, v := range nums {
		if pos, ok := numsMap[target-v]; ok && pos != k {
			return []int{k, pos}
		}
	}

	return make([]int, 2)
}
