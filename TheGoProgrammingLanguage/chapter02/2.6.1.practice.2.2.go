/*
练习2.2:写一个类似于cf的通用的单位转换程序，从命令行参数或者标准输入(如果没有参数)获取数字，
然后将每一个数字转换为以摄氏温度和华氏温度表示的温度，以英寸和米表示的长度单位，以磅和千克表示的重量，等等。
 */
package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"tgpl/chapter02/commconv"	//导入自己编写的包
)

func main() {
	/*从命令行读取*/
	if len(os.Args[1:]) > 0 {
		for _,arg := range os.Args[1:] {
			temp,err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "commconv error:%v\n", err)
				os.Exit(1)
			}
			Commconv(temp)
		}	
	} else {
		//使用bufio创建一个扫描器, 从 stdin 读取输入
		input := bufio.NewScanner(os.Stdin);
		for input.Scan() {
			temp,err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "commconv error:%v\n", err)
				os.Exit(1)
			}
			Commconv(temp)
		}
	}
}

func Commconv(temp float64) {
	ft := commconv.Fahrenheit(temp)	//类型显示转换
	ct := commconv.Celsius(temp)
	fmt.Printf("%s = %s, %s = %s\n", ft, commconv.FToC(ft), ct, commconv.CToF(ct))
	meter := commconv.Meter(temp)
	inch := commconv.Inch(temp)
	fmt.Printf("%s = %s, %s = %s\n", meter, commconv.MeterToInch(meter), inch, commconv.InchToMeter(inch))
	pound := commconv.Pound(temp)
	kg := commconv.Kilogram(temp)
	fmt.Printf("%s = %s, %s = %s\n", pound, commconv.PoundToKilogram(pound), kg, commconv.KilogramToPound(kg))
}