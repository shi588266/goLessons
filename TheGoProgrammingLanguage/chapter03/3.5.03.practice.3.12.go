/*练习 3.12:编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。*/
package main

import (
	"fmt"
	"strings"
)

func IsLikely(str1 string, str2 string) bool {
	str1Len := len(str1)
	str2Len := len(str2)
	if str1Len != str2Len {
		return false
	}
	/*遍历每一个字符,判断是否出现再另一个字符串中*/
	for i := 0; i < str1Len; i++ {
		if !strings.Contains(str2, str1[i:i+1]) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abcdefga"
	s2 := "gfedcba"
	if IsLikely(s1, s2) {
		fmt.Printf("'%s' and '%s' is accordingly but not same\n", s1, s2)
	} else {
		fmt.Printf("'%s' and '%s' are not accordingly\n", s1, s2)
	}
}