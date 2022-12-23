package main

import "fmt"

func getFullName() (string, string, int) {
	return "sulfan", "aidid", 182
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}
