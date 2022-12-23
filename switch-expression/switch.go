package main

import (
	"fmt"
)

func main() {
	// contoh pengecekan terhadap nama

	var nama string
	nama = "sulfan"

	switch nama {
	case "nap":
		fmt.Println("Hello nap")
	case "sulfan":
		fmt.Println("Hello sulfan")
	default:
		fmt.Println("Kenalan Dong")

	}

	// switch dengan short statement
	switch lenght := len(nama); lenght >= 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	lenght := len(nama)

	switch {
	case lenght > 10:
		fmt.Println("Nama Terlalu panjang")
	case lenght < 5:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
