package main

import "fmt"

// defer
//
//	func logging() {
//		fmt.Println("Selesai Memanggil Function")
//	}
//
//	func runApplication(value int) {
//		defer logging()
//		fmt.Println("Run application")
//		result := 10 / value
//		fmt.Println("Result", result)
//
// }
//
// Panic
func endApp() {
	//fmt.Println("End APP")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

//func runApp(error bool) {
//	defer endApp()
//	if error {
//		panic("APLIKASI ERROR")
//	}
//	fmt.Println("Aplikasi berjalan")
//}

// recover
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	//runApplication(10)
	runApp(true)
	fmt.Println("Check")
}
