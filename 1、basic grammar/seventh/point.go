package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)

	a := 100
	ptr = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

}
