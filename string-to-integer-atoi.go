//https://leetcode-cn.com/problems/string-to-integer-atoi/
package main

import "fmt"
import "math"

func main() {

	// s := "   -42"
	// s := "4193 with words"
	// s := "+1"
	// s := "+-2"
	// s := "9223372036854775808"
	// s := "2147483648"
	// s := "-2147483649"
	// s := "words and 987"
	// s := "-2147483647"
	// s := "2147483800"
	s := "0-1"
	// s := "2-1"
	fmt.Println(math.MaxInt32, myAtoi(s))
}

func myAtoi(str string) int {

	res := 0
	flag := 0
	space := true
	for i := 0; i < len(str); i++ {
		if str[i] == 32 && space {
			continue
		}
		space = false
		// if flag == 0 && (str[i] == 43 || str[i] == 45) && res == 0 {
		// 	flag = int(str[i])
		// 	continue
		// }
		if flag == 0 {
			flag = int(str[i])
			if (str[i] == 43 || str[i] == 45) && res == 0 {
				continue
			}
		}
		if str[i] < 48 || str[i] > 57 {
			break
		}
		if flag == 45 && res > 0 {
			res = -res
		}

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && (int(str[i])-48 > math.MaxInt32%10)) {
			res = math.MaxInt32
			break
		}

		if res < math.MinInt32/10 || (res == math.MinInt32/10 && (-(int(str[i]) - 48) < math.MinInt32%10)) {
			res = math.MinInt32
			break
		}
		// fmt.Println(res, -(int(str[i]) - 48), math.MaxInt32%10, math.MaxInt32/10)
		if flag == 45 {
			res = res*10 - (int(str[i]) - 48)
		} else {
			res = res*10 + (int(str[i]) - 48)
		}

	}

	return res
}
