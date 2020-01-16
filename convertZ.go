//https://leetcode-cn.com/problems/zigzag-conversion/
package main

import "fmt"
import "strings"

func main() {
	numRows := 3
	// s := "LEETCODEISHIRING"
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i, flag := 0, false //0向下 1向上

	for _, c := range s {
		res[i] = res[i] + string(c)

		if i == numRows-1 || i == 0 {
			flag = !flag
		}

		if flag {
			i++
		} else {
			i--
		}
	}
	// var resStr string
	// for _, str := range res {
	// 	resStr += str
	// }
	return strings.Join(res, "")
}
