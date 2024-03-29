//华氏温度转换摄氏温度
package commconv
import "fmt"
//类型定义
type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Meter float64
type Inch float64
type Pound float64
type Kilogram float64

//常量定义
const (
	AbsoluteZeroC Celsius 	= -273.15	//最低温度
	FreezingC Celsius 		= 0 		//冰点温度
	BoillingC Celsius 		= 100		//沸点温度
	AbsoluteZeroK Kelvin 	= 273.15			//开式绝对零度
)
const (
	Meter2Inch	= 39.3700787
	Pound2Kilogram	= 0.45359237
)

//函数声明
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
func (i Inch) String() string {
	return fmt.Sprintf("%ginch", i)
}
func (p Pound) String() string {
	return fmt.Sprintf("%gpound", p)
}
func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}