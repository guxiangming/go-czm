package main

import "fmt"

const Pi float64 = 3.1415
const zero = 0.0
const (
	size int64 = 1024
	eof        = -1
)
const u, v float32 = 0, 3   //u=0.0
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量

const (
	c0 = iota //0
	c1 = iota //1
	c2 = iota //2
)

func GetNumber() int {
	return 100
}

//go 实现枚举 是通过 const 定义体实现

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

// const num = GetNumber()

func main() {
	fmt.Println(Sunday, zero)
}
