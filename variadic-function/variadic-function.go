package main

import "fmt"

func sumAll(num ...int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}

func main() {
	ptotal := sumAll(12, 12, 123, 2)
	fmt.Println(ptotal)

	slice := []int{12, 23, 23, 41, 24}
	total := sumAll(slice...)
	fmt.Println(total)
}
