package main

import "fmt"

func main() {
	num := checkNumber(10)
	printNum(num)
}

func checkNumber(num int) []int {
	var n []int
	for i := 0; i < num; i++ {
		n = append(n, i)
	}
	return n
}

func printNum(nums []int) {
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, " is even.")
		} else {
			fmt.Println(num, " is odd.")
		}
	}
}
