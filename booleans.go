package main

import (
	"fmt"
	"runtime"
)

func Abs(x int) int  {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}




func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	fmt.Println(runtime.GOOS)


	if var1 := 10; var1 > 1 {
		fmt.Println("test")
	}
}




