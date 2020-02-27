package main

func main() {
	var slice [3]string = [3]string{"a", "b", "c"}
	q2 := slice[0:1]
	q2 = [4]int{1, 2, 3, 4}
	//直接创建切片
	mySlice := make([]int, 5)
	summer := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(summer); i++ {

	}
	for i, v := range summer {

	}

	var oldslice = make([]int, 5, 10)

	len(oldSlice)
	cap(oldSlice)

	newSlice := append(oldSlice, 1, 2, 3)
	slice1 := []int{1, 2, 3, 4, 5}
}
