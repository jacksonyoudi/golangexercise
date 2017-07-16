package main

import "fmt"

func main() {
	//for i := 1; i < 25; i++ {
	//	for j:= 1; j < i; j++ {
	//		fmt.Print("G")
	//	}
	//	fmt.Print("\n")
	//}


	a := "abcdefg"
	for index, value := range a {
		fmt.Println(index)
		fmt.Print("%c", value)
		fmt.Println("\n")
	}
}
