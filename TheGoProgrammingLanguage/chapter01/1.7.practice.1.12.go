/*
修改莉萨如服务器以通过URL参数读取参数值
执行方式:
	go build 1.7.practice.1.12.go
 */
package main

import (
    "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
	"log"
	"net/http"
	//"strconv"
)
/*声明包内变量*/
//返回color.Color结构体数组, 第一个元素代表黑色, 第二个元素代表绿色
var palette = []color.Color{color.Black,color.RGBA{0x00, 0xff, 0x00, 0xff}}
const (
	blackIndex = 0	//画板中的第一种颜色
	greenIndex = 1	//画板中的下一种颜色
)
const (
	cycles 	= 5	//完整的X振荡器变化个数
	res 	= 0.001	//角度分辨率
	size 	= 100	//图像画布包含[-size~+size]
	nframes = 64	//动画中的帧数
	delay	= 8 	//以10ms为单位的帧间延迟, 8*10 = 80ms,帧间延迟是80毫秒
)


func main() {
	//生成随机数
	rand.Seed(time.Now().UTC().UnixNano())
	//以 webserver 方式提供对外访问
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			/*接收并处理 Request 参数*/
			for key,val := range r.Form {
				if key == "cycles" {
					fmt.Println("key = %q\n", val)
				}
				if key == "nframes" {
					fmt.Println("key = %q\n", val)
				}
				if key == "size" {
					fmt.Println("key = %q\n", val)
				}
				if key == "delay" {
					fmt.Println("key = %q\n", val)
				}
			}
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

/*
创建位图图像, 将位图图像编码为GIF动画,
参数是图像输出地方,类型是 io.Writer
莉萨如图像
 */
func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 	//y振荡器的相对频率,相对于x振荡器
	//gif.GIF 结构体字面量, 创建一个动画, 这里先只给一个结构体成员, 也就是帧数, 动画总帧数是64
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 		//相位差
	/*每次迭代产生一个动画帧*/
	for i := 0; i < nframes; i++ {
		//创建 201x201 的画板
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//使用黑色和白色两种颜色填充画板
		img := image.NewPaletted(rect, palette)
		/*内层循环通过设置一些像素点为绿色使之产生新的图像*/
		for t := 0.0; t < cycles*2*math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}