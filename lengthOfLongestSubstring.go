package main

import "fmt"

//SlidingWindow
func main() {
	s := "jbpnbwwd"
	// s := "abcabcbb"
	// s := "bbbbb"
	// s := "pwwkew"
	// s := " "
	// s := ""
	// s := "dvdf"
	// s := "abba"
	m := map[uint8]int{}
	left := 0 //左侧窗口 [left,right)
	maxLen := 0

	for i := 0; i < len(s); i++ { //i就是右侧窗口
		if _, ok := m[s[i]]; ok {
			// fmt.Println(left, m[s[i]], i)
			left = max(left, m[s[i]]+1)
			// left = m[s[i]]
		}
		//i - left +1 //窗口尺寸
		maxLen = max(i-left+1, maxLen)
		m[s[i]] = i
	}
	fmt.Println(maxLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring() {
	// s := "abcabcbb"
	// s := "jbpnbwwd"
	// fmt.Printf("%T\n",s)

	// if _,ok :=  interface{}(s).(string);ok {
	// 	fmt.Println(s)
	// }
	// if _,ok :=  interface{}(s).(string);!ok || len(s) < 2 {
	//        fmt.Println("s len less 2")
	//    }

	// s := "bbbbb"
	// s := "pwwkew"
	// s:= " "
	s := ""
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }
	tmpMap := make(map[uint8]uint8)
	strLen := len(s)
	if strLen < 2 {
		fmt.Println(strLen)
	}
	// fmt.Println(strLen)

	maxLen := 0

	for j := 0; j < strLen; j++ {
		for i := j; i < strLen; i++ {
			if _, ok := tmpMap[s[i]]; ok {
				mapLen := len(tmpMap)
				// fmt.Println(j,i,tmpMap)
				if mapLen > maxLen {
					maxLen = mapLen
				}
				tmpMap = make(map[uint8]uint8)
				break
			} else {
				tmpMap[s[i]] = s[i]
			}
		}
	}
	fmt.Println(maxLen)
}
