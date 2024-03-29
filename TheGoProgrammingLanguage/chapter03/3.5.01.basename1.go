// basename 溢出路径部分和 .后缀
package main
import "os"
import "fmt"
func basename(s string) string {
	//找最后一个'/',从后往前找
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break;
		}
	}
	//找最后一个'.',保留'.'之前的所有内容, 从后往前找
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break;
		}
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