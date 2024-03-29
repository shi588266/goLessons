/*
演示http.Get()功能
运行方式:
	编译 go build 1.5.01.fetch.go
	命令行运行: ./1.5.01.fetch https://baidu.com
*/
package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	//从命令输入读取一个url
	for _,url := range os.Args[1:] {
		// http.Get() 请求没有出错的话, 会返回一个 响应结构, 相应结构的Body域包含了服务端响应的可读数据流
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 使用 ioutil 读取响应数据流
		body,err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}