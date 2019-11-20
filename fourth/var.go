package main

func GetName(userName, nickName string) {
	return "test", "czm"

}

func main() {
	var v1 int = 10
	var v2 = 10
	v3 := 10
	var (
		v4 bool
		v5 string
		v6 [4]int
	)
	var v7 struct {
		f float64
	}
	var v8 *int
	var v9 map[string]int   //map字典
	var v10 func(a int) int //返回整形
	// 需要注意的是，变量在声明之后，系统会自动将变量值初始化为对应类型的零值，比如上述 v1 的值为 0，v2 的值空字符串，v3 的值为 false，依次类推，我们打印上述变量的值，可以看到如下输出：
	// i, j = j, i
	_, nickName := GetName()
}
