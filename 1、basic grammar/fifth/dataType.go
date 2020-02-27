package main

//布尔类型 bool
// 整型：int8、byte、int16、int、uint、uintptr 等
// 复数类型：complex64、complex128
// 字符串：string
// 字符类型：rune
// 错误类型：error

//指针 pointer
//数组 array
//切片 silce
//字典 map
//通道 chan
//结构体 struct
//接口 interface

func main() {
	var v1 bool
	v1 = true
	v2 := (1 == 2)

	// 	var b bool
	// b = 1 // 编译错误
	// b = bool(1) // 编译错误
	// 	var int_value_1 int8
	// int_value_2 := 8   // int_value_2 将会被自动推导为 int 类型
	// int_value_1 = int_value_2  // 编译错误
	// 由此可见，所有比较运算符在比较的时候都会考虑进数据类型的因素，所以不需要类似 PHP 中 === 和 !== 这种严格比较运算符。
	// 注：++ 或 -- 只能出现在语句中，不能用于表达式，故不参与优先级判断。
}
