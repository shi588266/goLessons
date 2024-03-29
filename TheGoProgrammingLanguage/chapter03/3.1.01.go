package main

import "fmt"

func main() {
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2
	fmt.Printf("%08b\n", x)			//00100010
	fmt.Printf("%08b\n", y)			//00000110
	fmt.Printf("%08b\n", x & y)		//00000010
	fmt.Printf("%08b\n", x | y)		//00100110
	fmt.Printf("%08b\n", x ^ y)		//00100100
	fmt.Printf("%08b\n", x &^ y)	//00100000
}