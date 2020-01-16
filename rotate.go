//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	pos := len(nums) - k
	nums = append(nums[pos:], nums[:pos]...)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

// func rotate(nums []int, k int)  {
//     a = append(a[:i], a[i+1:]...)
// }