package main

import (
	"errors"
	"fmt"
)

func add(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b

}

func main() {
	x, y := 1, 2
	z := add(&x, &y)
	fmt.Println(x, y, z)
}

// func add(a, b *int) (int, error) {
//     if (*a < 0 || *b < 0) {
//         err := errors.New("只支持非负整数相加")
//         return 0, err
//     }
//     *a *= 2
//     *b *= 3
//     return *a + *b, nil
// }

// func main()  {
//     x, y := -1, 2
//     z, err := add(&x, &y)
//     if (err != nil) {
//         fmt.Println(err.Error())
//         return
//     }
//     fmt.Printf("add(%d, %d) = %d\n", x, y, z)
// }

func add2(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err = errors.New("只支持正数")
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
}
