// echo1 输出其命令行参数
// 运行方法: go run {$file} args...
package main
import (
	"fmt"
	"os"
)

func main() {
	//声明字符串变量,不显式初始化,会被隐式赋值默认初始值
	var s,sep string
	/*
	使用os从命令行获取参数(如果命令行参数数据量比较大的情况下,这种方式比较低效)
	os.Args 变量是一个字符串slice, os.Args[0] 表示命令(可执行程序)本身
	os.Args[1] 表示命令(可执行程序)参数列表的第一个参数, 依次类推...
	 */
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s);
}