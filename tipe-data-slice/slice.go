package main

import "fmt"

func main() {
	var months = [...]string{
		"Januar",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	slice2 := months[11:]
	fmt.Println(slice2)

	slice3 := append(slice2, "napx") // ini menggunakan array baru karena kapasitasnya diluar range array
	fmt.Println(slice3)

	slice3[1] = "Ubah Data" //maka ketika di ubah data array di var months tidak terubah
	fmt.Println(slice3)

	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "sulfan"
	newSlice[1] = "aidid"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	for _, i := range copySlice {
		if i == "sulfan" {
			fmt.Println("yeah")
		}
		fmt.Println(i)
	}

}
