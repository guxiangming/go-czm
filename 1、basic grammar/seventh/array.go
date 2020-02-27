package main

import "fmt"

func main() {
	var a = [3]byte{1, 12, 3} // 长度为8的数组，每个元素为一个字节
	var b []int
	c := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	for i, v := range a {
		fmt.Println(v)
	}

	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}

	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2)
		}
	}
}
