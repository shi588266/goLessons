/*4.7 修改函数reverse，来反转一个UTF-8编码字符串中的字符元素，传入参数是该字符串对应的字节slice类型（[]byte）。
是否可以做到不重新分配内存就实现该功能。*/
package main

import (
	"fmt"
	//"unicode"
	"unicode/utf8"
)

/*
来反转一个UTF-8编码字符串中的字符元素，传入参数是该字符串对应的字节slice类型（[]byte）
 */
func mbreverse(s []byte) []byte {
	res := []byte{} 	//用于存储结果
	l := len(s)		//字符串的长度
	for i := l; i > 0; {
		//每次处理最后一个 字符
		r,size := utf8.DecodeLastRune(s[0:i])
		res = append(res, []byte(string(r))...)
		i-=size
	}
	return res
}

func main() {
	b := []byte("北京欢迎您welcome")
    fmt.Printf("%s\n", b)
    fmt.Printf("%s\n", mbreverse(b))
}