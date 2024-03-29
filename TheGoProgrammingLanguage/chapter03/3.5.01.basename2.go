// basename 溢出路径部分和 .后缀
package main
import "os"
import "fmt"
import "strings"

func basename(s string) string {
	//找最后一个'/',从后往前找
	slash := strings.LastIndex(s, "/")	//找字符串中某个字符最后出现的位置,如果没找到返回-1
	s = s[slash+1:]
	//找最后一个'.',保留'.'之前的所有内容, 从后往前找
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		s = s[:dot-1]
	}
	
	return s
}


func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Fprintf(os.Stderr, "Error: No dir or filename parameter is provided.\n")
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		fmt.Println(basename(arg))
	}
}