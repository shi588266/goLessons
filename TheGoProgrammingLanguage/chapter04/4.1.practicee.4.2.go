/*练习 4.2: 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"hash"
	"crypto/sha256"
	"crypto/sha512"
)

/*编码函数
@param str string 	待编码的字符串
@param bits	int64	编码输出结果的位数
*/
func crypt(str string, bits int64) string {
	//从 hash 包声明 Hash类型 变量
	var hs hash.Hash
	switch bits {
		case 512:	
			hs = sha512.New()
		case 384:
			hs = sha512.New384()
		default:
			hs = sha256.New()
	}
	//类型强制转换成字节slice类型
	hs.Write([]byte(str))
	return string(hs.Sum(nil))
}

func main() {
	//从命令行接收参数
	if len(os.Args[1:]) <= 0 {
		fmt.Printf("Error:lack of args.Usage:%s 'sth tobe crypted' 'crypted bits'\n", os.Args[0])
		os.Exit(1)
	}
	bits := os.Args[1]
	str := os.Args[2]
	/*第1个参数是待转换的字符串, 第2个参数是转换进制, 第3个参数是转换结果存储的整型位数*/
	b,err := strconv.ParseInt(bits, 10, 64)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
	}
	fmt.Printf("%s crypto with algrithum sha%s is: %s\n", str, bits, crypt(str, b))
}

