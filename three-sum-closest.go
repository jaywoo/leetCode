//https://leetcode-cn.com/problems/3sum-closest/
package main

import "fmt"

// import "math"

func main() {
	// nums := []int{0, 1, 1, -3}
	// target := 1
	// nums := []int{-1, 2, 1, -4}
	// nums := []int{4, 5, 6}
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sortNums(nums)
	sum := nums[0] + nums[1] + nums[2]
	for k := 0; k < len(nums)-2; k++ {
		i := k + 1
		j := len(nums) - 1

		for i < j {
			tmpSum := nums[k] + nums[i] + nums[j]

			if tmpSum == target {
				return target
			}
			if abs(target-tmpSum) < abs(target-sum) {
				sum = tmpSum
			}
			if tmpSum < target {
				i++
			} else if tmpSum > target {
				j--
			}
		}
	}
	return sum
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func sortNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1
		for i < j {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
			j--
		}
	}
}
