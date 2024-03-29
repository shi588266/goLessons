/*
练习 1.8:修改 fetch 程序添加一个 http:// 前缀(假如该URL参数缺失协议前缀)。可能会用到strings.HasPrefix。
运行方式:
	编译 go build 1.5.01.fetch.go
	命令行运行: ./1.5.01.fetch https://baidu.com
*/
package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"strings"
)

func main() {
	//从命令输入读取一个url
	for _,url := range os.Args[1:] {
		//添加一个 https:// 前缀(假如该URL参数缺失协议前缀)
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		// http.Get() 请求没有出错的话, 会返回一个 响应结构, 相应结构的Body域包含了服务端响应的可读数据流
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 使用 io.Copy 读取响应数据流
		body,err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}