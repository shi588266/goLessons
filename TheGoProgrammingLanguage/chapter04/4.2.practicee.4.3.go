/*练习4.3:重写reverse函数，使用数组指针代替slice.*/
package main

import(
	"fmt"
)
/*翻转*/
func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

/*翻转数组指针*/
func reverse2(sp *[]int) {
	for i, j := 0, len(*sp)-1; i < j; i, j = i+1, j-1 {
		(*sp)[i], (*sp)[j] = (*sp)[j], (*sp)[i]
	}
}

func main() {
	arr := []int{1,2,3,4,5,6}
	as := arr[:]
	fmt.Println("arr:", arr)
	fmt.Println("as:", as)

	reverse2(&arr)
	fmt.Println("reverse arr:", arr)
	
	reverse(as)
	fmt.Println("reverse as:", arr)
}