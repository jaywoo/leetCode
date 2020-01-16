//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
package main

import (
	"fmt"
)

func main() {
	digits := "223"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	return iterStr(digits, 0)
}

func iterStr(digits string, currIndex int) []string {
	res := []string{}

	mapping := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	if len(digits) == currIndex+1 {
		for _, letter := range mapping[uint8(digits[currIndex])] {
			res = append(res, letter)
		}
		return res
	}

	for _, letter := range mapping[uint8(digits[currIndex])] {
		nextStrArray := iterStr(digits, currIndex+1)
		for _, str := range nextStrArray {
			res = append(res, letter+str)
		}
	}
	return res
}
