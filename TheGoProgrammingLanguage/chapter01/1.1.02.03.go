// echo3 使用strings 包的 join 函数
// 运行方式:go run 1.1.02.03.go testinput arg2
package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
	strings.Join 接收两个参数, 连个都是字符串, 返回两个字符串拼接之后的字符串
	 */
	fmt.Println(strings.Join(os.Args[1:], " "));	/*输出结果:testinput arg2*/
	fmt.Println(os.Args[1:]);	/*输出结果:[testinput arg2]*/
}