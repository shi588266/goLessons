//dup2 打印输入中多次出现的行的个数和文本// 它从 stdin 或指定的文件列表读取
// 运行方式:  go run .\1.3.01.dup2.go 然后输入几行内容, 按ctrl+z结束
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	//创建一个map, map是一个键名唯一的键值对集合
	counts := make(map[string]int)
	//从命令参数读取文件列表
	for _,fileName := range os.Args[1:] {
		//一次性读取文件内容到大块内存
		data,err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3%v\n", err)
			continue
		}
		//使用 换行符 分隔读取到的内容
		for _,lineText := range strings.Split(string(data), "\n") {
			counts[lineText]++
		}
	}
	//打印重复行
	for lineText,cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, lineText)
		}
	}
}

/*数组作为函数参数是引用传递,*/
func countLines(file *os.File, counts map[string]int) {
	//注意:忽略input.Err()中可能的错误
	input := bufio.NewScanner(file);
	for input.Scan() {
		counts[input.Text()]++;
	}
}