// 简易的web服务器
// go run 1.7.01.server2.go &
package main

import (
	"log"
	"fmt"
	"net/http"
	"sync"
)
var mu sync.Mutex
var cnt int

func main() {
	/*第2个参数是, 请求回调处理函数*/
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
请求回调处理函数
@params w http响应写入器
@params r http请求
 */
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	cnt++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", cnt)
	mu.Unlock()
}