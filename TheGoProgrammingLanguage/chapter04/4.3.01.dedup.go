/*使用map实现集合set功能,去除读取文件的重复行*/
package main

import (
	"fmt"
	"os"
	"bufio"
)

func dedup() {
	//创建一个map
	seen := make(map[string]bool)
	/*bufio从标准输入读取文件名*/
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	//处理读取异常
	if err := input.Err(); err != nil {
		fmt.Println("Read error")
	}
}

func main() {
	fmt.Println("start reading...")
	dedup()
}