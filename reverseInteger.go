//https://leetcode-cn.com/problems/reverse-integer/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1463847412))
	// fmt.Println(reverse(123))
}

func reverse(x int) int {
	// max, min := 0x7fffffff, 0x80000000
	// 2^31-1=2147483647,-2^31=-2147483648
	max, min := 2147483647, -2147483648
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > max/10 || res == max/10 && pop > max%10 {
			return 0
		}
		if res < min/10 || (res == min/10 && pop < min%10) {
			return 0
		}
		res = res*10 + pop

	}

	return res
}
