package main

import "fmt"
import "bytes"

func intToString(values []int) string {
	//声明缓冲变量, 变量类型是 bytes.Buffer
	var buf bytes.Buffer
	//WriteByte()方法向bytes.Buffer里追加文字符号, WriteRune()方法向bytes.Buffer里追加任意utf-8编码的文字符号
	buf.WriteByte('[')
	for i,val := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1,2,3,4,5,6}))
}