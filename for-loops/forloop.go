package main

import "fmt"

func main() {
	tolongCetak := 1

	for tolongCetak <= 10 { // selama kondisi belum true maka lakukan perulangan sampai kondisi false
		fmt.Println("Perulangan ke ", tolongCetak)
		tolongCetak++ // ditambah satu
	}

	for tolongCetak := 1; tolongCetak <= 10; tolongCetak++ {
		fmt.Println("Perulangan ke ", tolongCetak)
	}

	sliceA := []string{"Apel", "Semangka", "Pisang"}

	for i := 0; i < len(sliceA); i++ {
		fmt.Println(sliceA[i])
	}

	for _, buah := range sliceA {
		fmt.Println(buah)
	}

	person := make(map[string]string)
	person["Name"] = "Sulfan"
	person["Alamat"] = "Cendrawasih"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
