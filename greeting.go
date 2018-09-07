package main

import "fmt"

func main() {
	fmt.Println("test")
	greeting(1, 2, 3, 4, 5)
	fmt.Println("over")
}

func greeting(a ...int) {
	fmt.Println("hello,world!")
	for i, v := range a {
		fmt.Println(i, v)
	}
}
