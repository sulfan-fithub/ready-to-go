package main

import "fmt"

func calculateNum(num1, num2 int) int {
	result := num1 + num2
	return result
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	resultStr := getHello("Sulfan")
	fmt.Println(resultStr)
	fmt.Println(getHello(""))

	resultNum := calculateNum(20, 23)
	fmt.Println(resultNum)

}
