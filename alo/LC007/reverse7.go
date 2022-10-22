package main

import "fmt"

func main() {
	t := reverse(123)
	fmt.Println(t)
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		r := x % 10
		x = x / 10
		result = result*10 + r
	}
	if result > 1<<31-1 || result < -(1<<31) {
		return 0
	}
	return result
}
