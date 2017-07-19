package main

import "fmt"

func main()  {
	a := Add10()
	fmt.Println(a(1))
	b := Add100(10)(20)
	fmt.Println(b)
}

func Add10() func(b int) int {
	return func(b int) int {
		return b + 10
	}
}

func Add100(a int ) func(b int) int {
	return func(b int) int {
		return a + b
	}
}