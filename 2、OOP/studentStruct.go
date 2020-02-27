package main

// Go 语言中也不支持构造函数、析构函数，取而代之地，可以通过定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数：
type student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func Goodstudent(id uint, name string, male bool, score float64) *student {

	return &student{id: id, name: name, score: score}
}

func (s student) Getname() string {
	return s.name
}

func main() {

}
