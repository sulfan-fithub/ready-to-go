package main

import "fmt"

func main() {
	name := "nap"

	if name == "sulfan" {
		fmt.Println("Hello Sulfan")
	} else if name == "nap" {
		fmt.Println("Hello Nap")
	} else {
		fmt.Println("Bukan-Siapa-siapa")
	}

	sayHello("napx")

	if lenght := len(name); lenght > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}

func sayHello(param string) string {
	param = "HELL"
	return param
}
