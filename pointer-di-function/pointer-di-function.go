package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Malang", "Jawa Timur", "Malaysia"}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address) // ini tidak berubah
}
