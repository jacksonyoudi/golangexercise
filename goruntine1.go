package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait() // 进入线程， 不会阻塞
	go shortWait() // 不会干扰
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9) // 保证程序不会再 线程没有执行完就退出了
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
