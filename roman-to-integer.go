//https://leetcode-cn.com/problems/roman-to-integer/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("LX"))
}

func romanToInt(s string) int {
	mapping := make(map[string]int)
	mapping["I"] = 1
	mapping["V"] = 5
	mapping["X"] = 10
	mapping["L"] = 50
	mapping["C"] = 100
	mapping["D"] = 500
	mapping["M"] = 1000

	res := 0
	for k, v := range s {
		v := string(v)
		res += mapping[v]
		if k != 0 {
			pre := string(s[k-1])
			if mapping[v] > mapping[pre] {
				res -= mapping[pre] * 2
			}
		}

	}
	return res
}
