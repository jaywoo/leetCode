//https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/21/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1,1,2}
	removeDuplicates(nums)
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

}
//数据为有序数组
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	return i+1
}
