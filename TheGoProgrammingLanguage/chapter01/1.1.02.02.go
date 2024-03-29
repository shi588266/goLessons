// echo2 输出其命令行参数
package main
import (
	"fmt"
	"os"
)

func main() {
	//声明字符串变量,不显式初始化,会被隐式赋值默认初始值
	var s,sep string
	/*
	range可以迭代地从字符串或slice中获取数据, 每次迭代获取一双键值对
	Args[1:] 是省略了第二个截至索引值, 它获取从第1个参数到最后一个参数
	 */
	for key, value := range os.Args[1:] {
		s += sep + value
		sep = " "
		key += 1
	}
	fmt.Println(s);
}