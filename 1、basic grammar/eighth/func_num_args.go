package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sum(slice...)
}

func sum(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}
