package main

import (
	"fmt"
	"os"
	"strconv"
	"tgpl/chapter02/tempconv"	//导入自己编写的包
)

func main() {
	for _,arg := range os.Args[1:] {
		temp,err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}
		ft := tempconv.Fahrenheit(temp)	//类型显示转换
		ct := tempconv.Celsius(temp)
		fmt.Printf("%s = %s, %s = %s\n", ft, tempconv.FToC(ft), ct, tempconv.CToF(ct))
	}
}
