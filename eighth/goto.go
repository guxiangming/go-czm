package main

import "fmt"

func main() {

	arr := [][]int{{1}, {2, 3}, {4, 5}}
	fmt.Println(arr[2][0])
	goto EXIT
	fmt.Println(arr[2][1])
EXIT:
	fmt.Println(arr[0][0])

	fmt.Println(arr[1][0])
	fmt.Println(len(arr))
}
