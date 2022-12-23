package main

import "fmt"

func random() interface{} {
	return "OK"
}
func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value : ", value, "is string")
	case int:
		fmt.Println("Int", value, "is int")
	default:
		fmt.Println("Unoknown type")
	}

}
