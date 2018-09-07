package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// 声明并创建一个通道，通道类型是string

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	fmt.Printf("1")
	ch <- "Tripoli"
	fmt.Printf("1")
	ch <- "London"
	fmt.Printf("1")
	ch <- "Beijing"
	fmt.Printf("1")
	ch <- "Tokio"

	//传入值
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		// 取出值
		input = <-ch

		fmt.Printf("%s ", input)
		fmt.Printf("3")
	}
}
