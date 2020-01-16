//https://leetcode-cn.com/problems/4sum/
package main

import "fmt"
import "sort"

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 4 {
		return result
	}

	sort.Ints(nums)
	for s, num := range nums {
		newTarget := target - num
		newNums := nums[s:]
		for i := 0; i < len(newNums)-2; i++ {
			j := i + 1
			k := len(newNums) - 1
			for j < k {
				tmp := newTarget - newNums[i] + newNums[j] + newNums[k]

				if tmp == 0 {
					result = append(result, []int{num, newNums[i], newNums[j], newNums[k]})
					break
				}

				if tmp < 0 {
					j++
				} else {
					k--
				}

				for newNums[j] == newNums[j-1] {
					j++
				}
				for k < len(newNums) && newNums[k] == newNums[k+1] {
					k--
				}
			}
		}
	}
	return result
}
