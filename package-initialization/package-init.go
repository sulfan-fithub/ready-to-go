package main

import (
	"basic-in-go/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
