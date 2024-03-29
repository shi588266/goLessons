package commconv

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

func MeterToInch(m Meter) Inch {
	return Inch(m * Meter2Inch)
}
func InchToMeter(inch Inch) Meter {
	return Meter(inch / Meter2Inch)
}
func PoundToKilogram(p Pound) Kilogram {
	return Kilogram(p * Pound2Kilogram)
}
func KilogramToPound(kg Kilogram) Pound {
	return Pound(kg / Pound2Kilogram)
}