/*练习3.13:编写 KB 、MB的常量声明，然后扩展到YB*/
package main
// 10 * 100 		10 * 100 * 1000 		10 * 100 * 1000 * 1000 		10*100*一千的n次方
const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

