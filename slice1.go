package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	for k, v := range slice {
		fmt.Println(k)
		fmt.Println(v)
	}
}
