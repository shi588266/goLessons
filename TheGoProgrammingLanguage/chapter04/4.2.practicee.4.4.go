/*4.4：编写一个函数 rotate，实现一次遍历就可以完成元素旋转。*/
package main

import (
	"fmt"

)

/*
把slice向左旋转 pos 个位置
后面 len() - pos 个元素整体向左移动
前面 pos 个元素依次追加到最后
 */
func rotate(s []int, pos int) []int {

	rSlc := s[pos:]		//切片获得后面的slice
	for i := 0; i < pos; i++ {
		rSlc = append(rSlc, s[i])
	}
	return rSlc;
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := rotate(a, 3)
	fmt.Println(r)
}