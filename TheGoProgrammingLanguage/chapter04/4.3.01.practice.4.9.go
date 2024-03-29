/*练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
在第一次调用 Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。*/
package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	wordfreq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//这样可以按单词而不是按行输入
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordfreq[word]++
	}

	fmt.Printf("word\t count\n")
	for w, n := range wordfreq {
		fmt.Printf("%s\t%d\n", w, n)
	}
}