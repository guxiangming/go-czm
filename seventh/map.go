package main

import "fmt"

func main() {
	var testMap map[string]int
	testMap = map[string]int{"one": 1, "two": 2, "three": 3}
	key := "two"
	v, ok := testMap[key]
	fmt.Printf("The element of key %q:%d\n", ok, v)
	fmt.Println(testMap)
	testMap["four"] = 4
	fmt.Println(testMap)
	delete(testMap, "one")
	fmt.Println(testMap)
	key := make([]string, 0)
	for v, k := range testMap {
		key = append(key, k)
	}
	sort.String(key)
	//sort.Int()
}
