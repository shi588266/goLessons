/*练习 4.6: 编写一个函数，原地将一个UTF8编码的[]byte类型的slice中相邻的空格(参考unicode.IsSpace)替换成一个空格返回*/

package main

import(
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
DecodeRune(p []byte) (r rune, size int)
函数解码p开始位置的第一个utf-8编码的码值，返回该码值和编码的字节数。
 */
func rmvDupSpace(s []byte) []byte {
	var i int
	for l := 0; l < len(s); {
		/*读取一个字符,注意一个字符不一定是一个字节*/
		r, size := utf8.DecodeRune(s[i:])
		l += size
		/*判断是否是空字符*/
		if unicode.IsSpace(r) {
			//看上一个字节是否是空格
			if i > 0 && s[i-1] == byte(32) {
				//如果上一个字节是空格就跳过当前空格把后面的内容覆盖写入
				copy(s[i:], s[i+size:])
			} else {
				//如果上一个字节不是空格, 就把当前字节设置为空格,并把剩余的内容覆盖写入
				s[i] = byte(32)
				copy(s[i+1:], s[i+size:])
				i++
			}
		} else {
			//如果当前字符不是空格, 继续处理下一个字符即可
			i+=size
		}
	}
	return s[0:i]
}

func main() {
	var byteSlice = []byte{'h', 'e', 'l', 'l', 'o', ' ', ' ', 'w', 'o', 'r', 'l', 'd'}
	newBs := rmvDupSpace(byteSlice)
	fmt.Printf("%s\n",newBs)

	newBs2 := rmvDupSpace2(byteSlice)
	fmt.Printf("%s\n",newBs2)
}

func rmvDupSpace2(s []byte) []byte {
	prev := s[0]
	out := s[0:1]
	for _,b := range s[1:] {
		if b == prev && unicode.IsSpace(rune(prev)) {
			continue
		}
		out = append(out, b)
		prev = b
	}
	return out
}