
[TOC]

# chapter01
## 命令行参数  
### os.Args  
os 包提供一些函数和变量，以与平台无关的方式和操作系统打交道, 变量 os.Args 是一个字符串slice,它是一个动态容量的顺序数组s,可以通过s[1]来访问单个元素，通过s[m:n]来访问一段连续子区间，数组长度用 len(s)表示,所有的索引使用半开区间，即包含第一个索引，不包含最后一个索引,例如，slice s[m:n]，其中，0≤m≤n≤len(s)，包含n-m 个元素。
os.Args 的第一个元素是os.Args[e]，它是命令本身的名字;

### 变量  
如果变量没有明确地初始化，它将隐式地初始化为这个类型的空值

### range  
range产生一对值:索引和这个索引处元素的值
```go
func main() {
    //声明字符串变量,不显式初始化,会被隐式赋值默认初始值
    var s,sep string
    /*
    range可以迭代地从字符串或slice中获取数据, 每次迭代获取一双键值对
    Args[1:] 是省略了第二个截至索引值, 它获取从第1个参数到最后一个参数
     */
    for key, value := range os.Args[1:] {
        s += sep + value
        sep = " "
        key += 1
    }
    fmt.Println(s);
}
```

### strings.Join()  
strings.Join 接收两个参数, 连个都是字符串, 返回两个字符串拼接之后的字符串
```go
func main() {
    /*
    strings.Join 接收两个参数, 连个都是字符串, 返回两个字符串拼接之后的字符串
     */
    fmt.Println(strings.Join(os.Args[1:], " "));    /*输出结果:testinput arg2*/
    fmt.Println(os.Args[1:]);   /*输出结果:[testinput arg2]*/
}
```

### make() 
make 是内置的函数, 可以用来新建map，它还可以有其他用途.

### map()  
map 存储一个键名唯一的键/值对集合，并且提供常量时间的操作来存储、获取或测试集合中的某个元素;
因为要求键名唯一所以键名必须是其值能够进行相等(==)比较的任意类型,值可以是任意类型;
`map[string]int` 这个例子中，键的类型是字符串，值是 int;

```go
func main() {
    //创建一个map, map是一个键名唯一的键值对集合
    counts := make(map[string]int);
    //使用bufio创建一个扫描器, 从 stdin 读取输入
    input := bufio.NewScanner(os.Stdin);

    //循环读取
    for input.Scan() {
        counts[input.Text()]++;
    }
    //输出重复行
    for lineText,cnt := range counts {
        if cnt > 1 {
            fmt.Println(lineText);
        }
    }
}
code
```

### bufio 包  
bufio包，使用它可以简便和高效地处理输人和输出。其中一个最有用的特性是称为扫描器(Scanner)的类型，它可以读取输人，以行或者单词为单位断开，这是处理以行为单位的输人内容的最简单方式;
每一次调用input.Scan()读取下一行，并且将结尾的换行符去掉;通过调用input.Text()来获取读到的内容;

### os.Open()  
函数os.Open() 返回两个值。第一个是打开的文件(*os.File)，该文件随后被Scanner读取。第二个返回值是一个内置的error类型的值。如果err等于特殊的内置nil值，标准文件成功打开.
```go
func main() {
    //创建一个map, map是一个键名唯一的键值对集合
    counts := make(map[string]int)
    //从命令参数读取文件列表
    fileNameArray := os.Args[1:]
    if len(fileNameArray) <= 0 {
        //如果命令行没有指定文件列表, 就从标准输入读取
        countLines(os.Stdin, counts)
    } else {
        for _,fileName := range fileNameArray {
            file,err := os.Open(fileName)
            //是否读取错误
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2%v\n", err)
                continue
            }

            countLines(file, counts)
            file.Close()    //读取结束关闭文件指针
        }
    }
    //打印重复行
    for lineText,cnt := range counts {
        if cnt > 1 {
            fmt.Printf("%d\t%s\n", cnt, lineText)
        }
    }
}

/*数组作为函数参数是引用传递,*/
func countLines(file *os.File, counts map[string]int) {
    //注意:忽略input.Err()中可能的错误
    input := bufio.NewScanner(file);
    for input.Scan() {
        counts[input.Text()]++;
    }
}
```

### io.ReadFile()  
ReadFile 函数(从io/ioutil包)，它读取整个命名文件的内容

### strings.Split()
strings.Split 函数，它将一个字符串分割为一个由子串组成的slice.
实际上，bufio.Scanner、ioutil.ReadFile以及ioutil.WriteFile使用*os.File 中的Read和write 方法，但是大多数程序员很少需要直接访问底层的例程。像 bufio 和io/ioutil 包中上层的方法更易使用。

```go
func main() {
    //创建一个map, map是一个键名唯一的键值对集合
    counts := make(map[string]int)
    //从命令参数读取文件列表
    for _,fileName := range os.Args[1:] {
        //一次性读取文件内容到大块内存
        data,err := ioutil.ReadFile(fileName)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3%v\n", err)
            continue
        }
        //使用 换行符 分隔读取到的内容
        for _,lineText := range strings.Split(string(data), "\n") {
            counts[lineText]++
        }
    }
    //打印重复行
    for lineText,cnt := range counts {
        if cnt > 1 {
            fmt.Printf("%d\t%s\n", cnt, lineText)
        }
    }
}
```

### io.Copy(dst,src)  
从src 读，并且写人dst

### goroutine  
当一个goroutine 试图在一个通道上进行发送或接收操作时，它会阻塞，直到另一个goroutine 试图进行接收或发送操作才传递值，并开始处理两个goroutine.
```go
/**
 并发请求 url
 @params url string 要请求的url地址
 @params ch chan<- string, 是一个channel,通道里存储的内容是string
 */
func fetch(url string, chnl chan<- string) {
    start := time.Now()
    resp,err := http.Get(url)
    if err != nil {
        chnl <- fmt.Sprint(err) //把错误信息发送到channel
        return  //请求失败要立即返回
    }
    nbytes,err := io.Copy(ioutil.Discard, resp.Body)    //把内容写入到ioutil.Discard 进行丢弃, 返回写入成功的字节数
    resp.Body.Close()   //防止资源泄露
    if err != nil {
        chnl <- fmt.Sprintf("while reading %s: %v", url, err)
        return  //写入失败要立即返回
    }
    /*获取响应数据响应时间并通过channel 返回*/
    seconds := time.Since(start).Seconds()
    chnl <- fmt.Sprintf("%.4fs %7d %s", seconds, nbytes, url)
}
```

# chapter02  
## 整型  
对于算术和逻辑(不含移位)的二元运算符，其操作数的类型必须相同;
于每种类型 T，若允许转换，操作T(x) 会将x的值转换成类型 t;
缩减大小的整型转换，以及整型与浮点型的相互转换，可能改变值或损失精度;
浮点型转成整型，会舍弃小数部分，``趋零截尾``(正值向下取整，负值向上取整)

## 字符串  
### bytes  
bytes包为高效处理字节slice提供了Buffer 类型。Buffer 起初为空，其大小随着各种类型数据的写人而增长  
```go
package main

import "fmt"
import "bytes"

func intToString(values []int) string {
    //声明缓冲变量, 变量类型是 bytes.Buffer
    var buf bytes.Buffer
    //WriteByte()方法向bytes.Buffer里追加文字符号, WriteRune()方法向bytes.Buffer里追加任意utf-8编码的文字符号
    buf.WriteByte('[')
    for i,val := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", val)
    }
    buf.WriteByte(']')
    return buf.String()
}
//打印 [1, 2, 3, 4, 5, 6]
func main() {
    fmt.Println(intToString([]int{1,2,3,4,5,6}))
}
```

### 常量  
#### iota 常量生成器  
用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一  
```go
type Weekday int
//周日将对应0，周一为1，如此等等
const(
    Sunday Weekday = iota
    Monday Weekday
    Tuesday Weekday
    Wednesday Weekday
    Thursday Weekday
    Friday Weekday
    Saturday Weekday
)
```
下面是来自net包的例子，用于给一个无符号整数的最低5bit的每个bit指定一个名字:
```go
type Flags uint 
const(
    FlagUp Flags = 1 << iota    //is Up                  1
    FlagBraodcast               //支持广播访问           10
    FlagLoopback                //是回环接口            100
    FlagPointToPoint            //属于点对点链路       1000
    FlagMulticast               //支持多路广播访问     10000
)
```
>随着`iota`的递增，每个常量对应表达式`1<<iota`，是连续的2的幂，分别对应一个bit位置。使用这些常量可以用于测试、设置或清除对应的bit位的值;

```go
func IsUp(v Flags) bool {
    return v & FlagUp == FlagUp
}
func TurnDown(v *Flags) bool {
    *v &^= FlagUp
}
func SetBroadcast(v *Flags) bool {
    *v |= FlagBraodcast
}
func IsCast(v Flags) bool {
    return v&(FlagBraodcast|FlagMulticast) != 0
}
func main(){
    var v Flags=FlagMulticast | FlagUp
    fmt.Printf("%b %t\n",v,IsUp(v))                 // “10001 true"
    TurnDown(&v)fmt.Printf("%b %t\n",v,IsUp(v))     //“10000 false"
    SetBroadcast(&v)fmt.Printf("%b %t\n",v,IsUp(v)) //"10010 false"
    fmt.Printf("%b %t\n",v,IsCast(v))               //"10010 true
}
```

----
# chapter 04  
## 数组  
<mark> **数组的长度是数组类型的一个组成部分** </mark>，因此[3]int和[4]int是两种不同的数组类型。
数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我们可以直接通过==比较运算符来比较两个数组，只有当两个数组的每个索引对应的元素都是相等的时候数组才是相等的。
```go
a :=[2]int{1，2}
b :=「...]int{1，2}
c:=[2]int{1，3}
fmt.Println(a==b,a==c,b==c)// "true false false"
d :=[3]int{1，2}
//元素个数不同的数组是两种不同的类型, 无法进行比较
fmt.Println(a==d)//compile error: cannot compare [2]int == [3]int
```

## Slice
一个sice由三个部分构成:指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是sice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目:长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。

### len()返回slice的长度

### cap()返回slice的容量

### slice的切片操作`s[i:j]`  
slice的切片操作`s[i:j]`，其中`0≤i<j≤ cap(s)`，用于创建一个新的slice，引用s的从第i个元素开始到第i-1个元素的子序列。新的slice将只有i-i个元素。如果i位置的索引被省略的话将使用0代替，如果j被省略将使用 len(s) 代替;  
如果切片操作超出cap(s)的上限将导致一个 panic 异常，但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大.  

注意:  **slice执行切片得到的是一个新的slice, 但是这个新的slice仍旧由一个指针指向底层的数组, 等于新的slice在数据上和原来的slice是重叠的**

因为slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。换句话说，复制一个sice只是对底层的数组创建了一个新的slice别名
```go
//reverse reverses aslice of ints in place.
func reverse(s int){
    for i,j:=0,len(s)-1;i<j; i,j= i+1,j-1 {
        s[i]，s[j]= s[j],s[i]
    }
}
```
slice唯一合法的比较操作是和nil比较:``if slice == nil {/* ... */}``;
一个零值的slice等于nil。一个nil值的slice并没有底层数组。一个nil值的slice的长度和容量都是0

### 通过make来创建切片
```go
var slice []type = make([]type, len)
    slice  := make([]type, len)
    slice  := make([]type, len, cap)
```

### append（切片追加）
```go
var x[]int
x= append(x，1)
x=append(x，2，3)
x=append(x，4，5，6)
x=append(x，x...)        //append the slice x
fmt .Println(x)
//"[12345 612 3 4 5 6]"
```

## map 
一个map就是一个哈希表的引用,map类型可以写为map[K]V，其中K和V分别对应 key和value,
其中K对应的key必须是支持==比较运算符的数据类型;
禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效;
### map定义
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
```go
ages := map[string]int{"alice":31, "charlie":34} 
// equivlent
ages := make(map[string]int) 
ages["alice"] = 31 ages["charlie"] = 34 
```

Go语言中有个判断map中键是否存在的特殊写法；第二个是一个布尔值，用于报告元素是否真的存在。
```go
age, ok := ages["bob"] 
if !ok { 
    /* "bob" is not a key in this map; age == 0. */ 
} 
```
### map的遍历  
```go
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    scoreMap["王五"] = 60
    for k, v := range scoreMap {
        fmt.Println(k, v)
    }
}
```
只想遍历key的时候，可以按下面的写法：
```go
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    scoreMap["王五"] = 60
    for k := range scoreMap {
        fmt.Println(k)
    }
}
```
>此时将这种忽略value的map当作一个字符串集合  

### 元素为map类型的切片
```go
func main() {
    var mapSlice = make([]map[string]string, 3)
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
    fmt.Println("after init")
    // 对切片中的map元素进行初始化
    mapSlice[0] = make(map[string]string, 10)
    mapSlice[0]["name"] = "王五"
    mapSlice[0]["password"] = "123456"
    mapSlice[0]["address"] = "红旗大街"
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
}
```
### 值为切片类型的map  
```go
func main() {
    var sliceMap = make(map[string][]string, 3)
    fmt.Println(sliceMap)
    fmt.Println("after init")
    key := "中国"
    value, ok := sliceMap[key]
    if !ok {
        value = make([]string, 0, 2)
    }
    value = append(value, "北京", "上海")
    sliceMap[key] = value
    fmt.Println(sliceMap)
}
```

