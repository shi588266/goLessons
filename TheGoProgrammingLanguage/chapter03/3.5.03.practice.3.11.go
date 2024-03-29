/*练习3.11:完善comma函数，以支持浮点数处理和一个可选的正负号的处理。*/
package main
import "os"
import "fmt"
import "bytes"
import "strings"

func comma(s string) string {
	var buf bytes.Buffer	//创建 字节buffer类型 变量
	//看是否是有符号的数字字符串
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start++
	}
	/*看是否是否点数,小数点是否存在*/
	dotIdx := strings.Index(s, ".")
	if dotIdx != -1 {
		//处理整数部分, 字符下标从start到 dotIdx
		trans(s[start:dotIdx], &buf)
		//把 点号 追加进去
		buf.WriteRune('.')
		//处理小数点后面的部分, 按照格式要求,小数点后面的部分不需要使用逗号分隔
		buf.WriteString(s[dotIdx+1:])
	} else {
		trans(s[start:], &buf)
	}
	return buf.String()
}

/*
把字符串的内容, 转移到buf中, 以字符串最后一个字符为基准, 每3个为1组, 用逗号分隔
因为要修改buf的内容,所以第二个参数需要传指向 buf 的指针
 */
func trans(str string, buf *bytes.Buffer) {
	lnth := len(str)
	if len(str) < 3 {
		(*buf).WriteString(str[:])
		return
	}
	(*buf).WriteString(str[0:lnth%3])
	for i := lnth%3; i < lnth; i+=3 {
		if i != 0 {
			(*buf).WriteByte(',')
		}
		(*buf).WriteString(str[i:i+3])
	}
}

func main() {
	var str string = "-12345698.96314785"
	fmt.Printf("%s transfer to %s\n", str, comma(str))
	for _,arg := range os.Args[1:] {
		fmt.Printf("%s transfer to %s\n", arg, comma(arg))
	}
}