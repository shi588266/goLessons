/*unicode 码点字符数量统计*/
package main

import(
	"os"
	"fmt"
	"io"
	"bufio"
	"unicode"
	"unicode/utf8"
)

func main() {
	
	cnts := make(map[rune]int)			//用以统计Unicode字符出现数量
	var utflen [utf8.UTFMax + 1]int 	//用以统计不同码点的长度的次数(utf8编码的码点的长度最小1个字节最大4个字节)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()	//读取一个码点, 返回 rune,nbytes,error
		if err == io.EOF {	//读取结束
			break;
		}
		if err != nil {
			fmt.Fprintf(os.Stderr,"charcount:%v\n",err)
			os.Exit(1)
		}
		//读取到 替换字符 (无效字符会被替换为固定的码值)
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		cnts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c,n := range cnts{
		fmt.Printf("%q\t%d\n",c,n)
	}
	fmt.Print("\nlen\tcount\n")
	for i,n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n",i,n)
		}
	}
	if invalid > 0{
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}