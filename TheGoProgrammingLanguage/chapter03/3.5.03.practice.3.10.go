/*练习 3.10:编写一个非递归版本的comma函数，使用bytes.Bufer代替字符串链接操作。*/
package main
import "bufio"
import "fmt"
import "os"
import "bytes"

func comma(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	} 
	var buf bytes.Buffer
	buf.WriteString(s[:length%3])	//如果字符串除以3有余数,先写入余数长度的字符
	//然后每次读取一段,也就是3个数字字符, 当然读入每一段之前要先写入一个逗号
	for i := length%3; i < length; i+=3 {
		// 如果整除3, 不用先写入逗号
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i:i+3])
	}
	//最后转换为字符串
	return buf.String()
}

func main() {
	//从命令行接收输入
	if len(os.Args[1:]) > 0 {
		for _,arg := range os.Args[1:] {
			fmt.Printf("%s transfer to %s\n", arg, comma(arg))
		}
	} else {	//如果没有命令行参数,就提示用户从终端输入
		fmt.Fprintf(os.Stdout, "please enter an integer\n")
		//使用 bufio 创建一个扫描器, 从 stdin 读取输入
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			fmt.Printf("%s transfer to %s\n", input.Text(), comma(input.Text()))
		}
	}
}