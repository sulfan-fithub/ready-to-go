package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
	Married bool
}

func main() {
	var nap Customer
	nap.Name = "Sulfan Aidid"
	nap.Address = "Cendrawasih"
	nap.Age = 24
	nap.Married = false

	fmt.Println(nap)
	fmt.Println(nap.Name)
	fmt.Println(nap.Address)
	fmt.Println(nap.Age)
	fmt.Println(nap.Married)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
		Married: false,
	}
	fmt.Println(joko)
	fmt.Println("NAMA : ", joko.Name)
	fmt.Println(joko.Address)
	fmt.Println(joko.Age)
	fmt.Println(joko.Married)

	cali := Customer{"Cale", "Indonesia", 24, false}
	fmt.Println(cali)

	alig := Customer{
		Name:    "Aligman",
		Address: "Makassar",
		Age:     10,
		Married: true,
	}
	fmt.Println(alig)

}
