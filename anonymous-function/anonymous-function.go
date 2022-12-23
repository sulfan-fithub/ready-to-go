package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blackLiskRoot(name string) bool { //tanpa harus membuat function seperti ini
//	return name == "root"
//}

func main() {
	blacklist := func(name string) bool { // jadi bisa langsung membuat function dan dimasukkan kedalam variable
		return name == "admin"
	}

	registerUser("admin", blacklist) //
	registerUser("nap", blacklist)

	registerUser("root", func(name string) bool { // atau masukkan function kedalam kedalam parameter secara
		// langsung tanpa harus membuat nama function terlebih dahulu
		return name == "root"
	})

	registerUser("nap", func(name string) bool {
		return name == "root"
	})
}
