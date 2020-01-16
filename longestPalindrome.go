//https://leetcode-cn.com/problems/longest-palindromic-substring/
package main

import "fmt"

func main() {
	// s := "babad"
	// s := "cbbd"
	// s := "bb"
	// s := ""
	// if s == "" || len(s) < 2 {
	// 	return ""
	// }
	fmt.Println(s)
	start := 0
	end := 0
	strLen := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		tmpLen := max(len1, len2)
		if strLen < tmpLen {
			strLen = tmpLen
			start = i - (strLen-1)/2
			end = i + strLen/2
		}
	}

	fmt.Println(s[start : end+1])
}

//方法四：中心扩展算法 ,只看懂这一种方法了
func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left = left - 1
		right = right + 1
	}
	return right - left - 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
