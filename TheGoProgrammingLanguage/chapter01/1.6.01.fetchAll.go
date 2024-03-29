// fetchall 并发获取URL并报告它们的时间和大小
// ./test https://baidu.com https://alibaba.com https://www.qq.com/
package main

import (
	"os"
	"io"
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
)
/**
 并发请求 url
 @params url string 要请求的url地址
 @params ch chan<- string, 是一个channel,通道里存储的内容是string
 */
func fetch(url string, chnl chan<- string) {
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil {
		chnl <- fmt.Sprint(err)	//把错误信息发送到channel
		return	//请求失败要立即返回
	}
	nbytes,err := io.Copy(ioutil.Discard, resp.Body)	//把内容写入到ioutil.Discard 进行丢弃, 返回写入成功的字节数
	resp.Body.Close()	//防止资源泄露
	if err != nil {
		chnl <- fmt.Sprintf("while reading %s: %v", url, err)
		return	//写入失败要立即返回
	}
	/*获取响应数据响应时间并通过channel 返回*/
	seconds := time.Since(start).Seconds()
	chnl <- fmt.Sprintf("%.4fs %7d %s", seconds, nbytes, url)
}

func main() {
	//起始时间
	start := time.Now()
	/*使用内置make()创建channel, 注意这个语句,chan是关键字*/
	chnl := make(chan string)

	for _,url := range os.Args[1:] {
		go fetch(url, chnl)		/*go关键字启动一个 goroutine*/
	}
	/*从通道接收*/
	for range os.Args[1:] {
		fmt.Println(<-chnl)
	}
	//统计总耗时
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}