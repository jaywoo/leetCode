//https://leetcode-cn.com/problems/container-with-most-water/
package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea2Point(height))
}

//双指针
func maxArea2Point(height []int) int {
	area := 0
	h := 0
	lastHeight := 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}
		if h > lastHeight {
			w := j - i + 1
			tmp := w * h
			if area < tmp {
				area = tmp
			}
			h = lastHeight
		}
	}

	return area
}

// 暴力
func maxArea(height []int) int {
	area := 0
	length := len(height)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			h := min(height[i], height[j])
			w := j - i
			tmp := w * h
			if area < tmp {
				area = tmp
			}
		}
	}
	return area
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
