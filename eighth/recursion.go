package main

const MAX = 50

var fibs [MAX]int

func main() {

}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
