//https://leetcode-cn.com/problems/integer-to-roman/
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(9/5, 9%5) //1,4
	fmt.Println(intToRoman(99))
}

func intToRoman(num int) string {
	if num < 1 || num > 39999 {
		return ""
	}

	base := 1
	mapping := make(map[int]string)
	mapping[1] = "I"
	mapping[5] = "V"
	mapping[10] = "X"
	mapping[50] = "L"
	mapping[100] = "C"
	mapping[500] = "D"
	mapping[1000] = "M"

	res := ""
	for num > 0 {
		pop := num % 10
		num /= 10

		rem := pop % 5
		mod := pop / 5

		//处理进位（4、9）
		if rem == 4 {
			if mod == 0 {
				res = mapping[base] + mapping[5*base] + res
			} else {
				res = mapping[base] + mapping[base*10] + res
			}
		} else if pop == 5 {
			res = mapping[5*base] + res
		} else {
			tmp := ""
			for i := 1; i <= rem; i++ {
				tmp += mapping[base]
			}

			if mod == 1 {
				res = mapping[5*base] + tmp + res
			} else {
				res = tmp + res
			}
		}
		base *= 10
	}
	return res
}
