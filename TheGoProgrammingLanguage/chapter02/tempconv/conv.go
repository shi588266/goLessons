package tempconv

//CToF 把摄氏温度转化为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32);
}
//FToC 把摄氏温度转化为华氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)/9*5);
}
//CToK 把摄氏温度转化为开氏温度
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC);
}
//KToC 把摄氏温度转化为开氏温度
func KToC(k Kelvin) Celsius {
	return Celsius(k + AbsoluteZeroK);
}