// https://leetcode-cn.com/problems/remove-element/
package main

import "fmt"

func main() {
	nums := []int{2}
	val := 3
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	fmt.Println(n, nums)
}

func removeElement1() {
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// val := 2
	// nums := []int{3, 2, 2, 3}
	nums := []int{2}
	val := 3
	j := len(nums)
	for i := 0; i < len(nums); i++ {
		if j == i {
			break
		}

		if nums[i] == val {
			for j > i {
				if nums[j-1] != val {
					nums[i] = nums[j-1]
					nums[j-1] = val
					break
				}
				j--
			}
		}
	}

	fmt.Println(j, nums)
}
