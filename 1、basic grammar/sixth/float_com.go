// 对于浮点类型需要被自动推导的变量，其类型将被自动设置为 float64，而不管赋值给它的数字是否是用 32 位长度表示的。因此，对于以上的例子，下面的赋值将导致编译错误：
// loat_value_1 = float_value_2  // float_value_2 是 float64 类型

package main

func main() {
	p := 0.000001
	// z=a+bi
	// 在 Go 语言中，复数支持两种类型：complex64（32位实部和虚部） 和 complex128（64位实部与虚部），对应的表示示例如下，和数学概念中的复数表示形式一致：
	var complex_value_1 complex64
	complex_value_1 = 1.10 + 10i
	complex_value_2 := 1.10 + 10i
	complex_value_3 := complex(1.10, 10)
}
