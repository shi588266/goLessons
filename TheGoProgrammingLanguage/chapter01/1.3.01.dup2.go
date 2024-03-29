//dup2 打印输入中多次出现的行的个数和文本// 它从 stdin 或指定的文件列表读取
// 运行方式:  go run .\1.3.01.dup2.go 然后输入几行内容, 按ctrl+z结束
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个map, map是一个键名唯一的键值对集合
	counts := make(map[string]int)
	//从命令参数读取文件列表
	fileNameArray := os.Args[1:]
	if len(fileNameArray) <= 0 {
		//如果命令行没有指定文件列表, 就从标准输入读取
		countLines(os.Stdin, counts)
	} else {
		for _,fileName := range fileNameArray {
			file,err := os.Open(fileName)
			//是否读取错误
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2%v\n", err)
				continue
			}

			countLines(file, counts)
			file.Close()	//读取结束关闭文件指针
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