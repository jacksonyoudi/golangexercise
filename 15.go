package main

import "fmt"

func main() {
	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)
}
