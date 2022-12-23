package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	var address2 *Address = &address1
	var databebas *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(databebas)

	var address4 *Address = new(Address)
	address4.Country = "Indoensia"
	fmt.Println(address4)
}
