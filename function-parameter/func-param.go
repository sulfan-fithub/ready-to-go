package main

import "fmt"

func sayHelloTo(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Sulfan", "aidid")
}
