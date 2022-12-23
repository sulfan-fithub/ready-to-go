package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Sulfan",
		"address": "Cendrawasih",
	}

	person["tittle"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"]) // ini cara memanggil Map
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Sulfan"

}
