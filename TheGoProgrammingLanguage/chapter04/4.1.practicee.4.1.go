/*练习 4.1:编写一个函数，计算两个SHA256哈希码中不同bit的数目。(参考2.6.2节的PopCount函数。)*/
package popcount
import "fmt"
import "crypto/sha256"

// pc[i] 是i的种群统计
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
//PopCount 返回x的种群统计
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + 
		pc[byte(x>>(1*8))] + 
		pc[byte(x>>(2*8))] + 
		pc[byte(x>>(3*8))] + 
		pc[byte(x>>(4*8))] + 
		pc[byte(x>>(5*8))] + 
		pc[byte(x>>(6*8))] + 
		pc[byte(x>>(7*8))])
}

func main() {
	//求一个任意字节 slice 类型的数据的hash摘要, 生成的摘要有256个bit位大小
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	PopCount()
}
