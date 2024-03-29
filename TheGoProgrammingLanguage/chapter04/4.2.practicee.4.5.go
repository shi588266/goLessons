/*练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。*/

package main

import(
	"fmt"
)

/*
消除切片中相邻重复的字符串
 */
func remvDup(ss []string) []string {
	lastStr := ss[0]	//第一个元素
	out := ss[0:1]	//切片的一部分

	for i,str := range ss {
		if i != 0 && str != lastStr {
			out = append(out, str)
		}
		lastStr = ss[i]
	}
	return out
}

func main() {
	data := []string{"one", "one", "one", "two", "two", "three", "three"}
	data = remvDup(data)
	fmt.Printf("%q\n", data)
}