package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Sulfan"
	middleName = "Al"
	lastName = "Aidid"

	return
}

func main() {
	PfirstName, PmiddleName, PlastName := getCompleteName()
	fmt.Println(PfirstName, PmiddleName, PlastName)
}
