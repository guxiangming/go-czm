package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	str = "hello world"
	ch := str[0]
	fmt.Printf("the lenght of %s is %d", str, len(str))
	str = str + "czm"
	str += "czm1121"
	str = "hello world"
	len := len(str)
	for i := 0; i < len; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
	for i, ch := range str {
		fmt.Println(i, ch)
	}

	v1 := uint(16)   // 初始化 v1 类型为 unit
	v2 := int8(v1)   // 将 v1 转化为 int8 类型并赋值给 v2
	v3 := uint16(v2) // 将 v2 转化为 uint16 类型并赋值给 v3
	v1 := "100"
	v2, err := strconv.Atoi(v1) //将字符串转换为整型
	v3 := 100
	v4 := strconv.Itoa(v3) //整形转字符串
	// strconv.ParseBool 字符串转换布尔
	// strconv,FormatBool 布尔值转换字符串
	// strconv.ParseUint() //无符号整型
	// strconv.FormatUint  整型转字符串
	// strconv.ParseFloat() 字符串浮点
	// strconv.FormatFloat 浮点字符串
	// strconv.Quote 字符串加引号
	// strconv.QuoteToASCII 字符串转换为
}
