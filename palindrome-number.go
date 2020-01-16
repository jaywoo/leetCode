//https://leetcode-cn.com/problems/palindrome-number/
package main

import "fmt"
import "math"

func main() {
	fmt.Println(isPalindrome(2222222222))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origX := x
	newX := 0
	for x > 0 {
		pop := x % 10
		x = x / 10
		if newX >= math.MaxInt32/10 {
			return false
		}
		newX = newX*10 + pop
	}

	return newX == origX
}
