//https://leetcode-cn.com/problems/longest-common-prefix/
package main

import "fmt"

func main() {
	// strs := []string{"dog", "racecar", "car"}
	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"aa", "a"}
	// strs := []string{"b", "a"}
	// strs := []string{"aaa", "aa", ""}
	// strs := []string{"a", "ac"}
	strs := []string{""}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	strLen := len(strs)
	if strLen < 1 {
		return ""
	}
	if strLen == 1 {
		return strs[0]
	}
	ans := strs[0]
	endPos := 0

	for _, s := range strs[1:] {
		if len(s) == 0 {
			return ""
		}
		sLen := len(s) - 1
		aLen := len(ans) - 1

		for k, c := range ans {

			if (k == sLen || k == aLen) && rune(c) == rune(s[k]) {
				endPos = k + 1
				break
			}
			if rune(c) != rune(s[k]) {
				endPos = k
				break
			}
		}
		ans = ans[:endPos]

	}
	return ans
}
