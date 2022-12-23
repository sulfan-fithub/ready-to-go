package main

import "fmt"

type Costumer struct {
	Name string
	Age  int
}

type Admin struct {
	Name string
	Role string
}

func (customer Costumer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name, "umur saya", customer.Age)
}

func (a Admin) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", a.Name, "Role Saya", a.Role)
}

func main() {
	nap := Costumer{Name: "Nap", Age: 25}
	eci := Admin{Name: "Eci", Role: "Admin"}

	nap.sayHello("Cali")
	eci.sayHi("Sita")
}
