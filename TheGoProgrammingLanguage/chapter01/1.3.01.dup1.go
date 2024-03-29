// dup1 输出标准输入中出现次数大于1的行，前面是次数
// 运行方式:  go run .\1.3.01.dup1.go 然后输入几行内容, 按ctrl+z结束
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个map, map是一个键名唯一的键值对集合
	counts := make(map[string]int);
	//使用  bufio 创建一个扫描器, 从 stdin 读取输入
	input := bufio.NewScanner(os.Stdin);

	//循环读取, 每次读取一行内容
	for input.Scan() {
		counts[input.Text()]++;
	}
	//输出重复行
	for lineText,cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d times ", cnt);
			fmt.Println(lineText + " accuors");
		}
	}
}