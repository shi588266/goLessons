// 简易的web服务器
// go run 1.7.01.server1.go &
package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	/*第2个参数是, 请求回调处理函数*/
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
请求回调处理函数
@params w http响应写入器
@params r http请求
 */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}