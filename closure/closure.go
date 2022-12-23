package main

import "fmt"

func main() {
	name := "Nap"
	counter := 0

	increment := func() {
		name = "sulfan"
		fmt.Println("Increment")
		counter++
	}

	increment()
	fmt.Println(name)

	fmt.Println(counter)
}
