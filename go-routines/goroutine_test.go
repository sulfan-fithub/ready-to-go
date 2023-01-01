package go_routines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("UPS")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(4 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "nap"
	channel <- "aidid"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

//func TestSelectChannel(t *testing.T) {
//	channel1 := make(chan string)
//	channel2 := make(chan string)
//
//	defer close(channel1)
//	defer close(channel2)
//
//	go GiveMeResponse(channel1)
//	go GiveMeResponse(channel2)
//
//	counter := 0
//	for {
//		select {
//		case data := <-channel1:
//			fmt.Println("Data dari channel 1", data)
//			counter++
//		case data := <-channel2:
//			fmt.Println("Data dari channel 2", data)
//			counter++
//		}
//
//		if counter == 2 {
//			break
//		}
//	}
//}

//func TestSelectChannel(t *testing.T) {
//	channel1 := make(chan string)
//	channel2 := make(chan string)
//
//	defer close(channel1)
//	defer close(channel2)
//
//	go GiveMeResponse(channel1)
//	go GiveMeResponse(channel2)
//
//	counter := 0
//	for {
//		select {
//		case data := <-channel1:
//			fmt.Println("Data dari channel 1", data)
//			counter++
//		case data := <-channel2:
//			fmt.Println("Data dari channel 2", data)
//			counter++
//		default:
//			fmt.Println("Menunggu Data")
//		}
//
//		if counter == 2 {
//			break
//		}
//	}
//}
