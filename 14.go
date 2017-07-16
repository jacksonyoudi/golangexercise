package main

import "fmt"

type Interger int

func (a Interger) Less(b Interger) bool {
	return a < b
}

func main() {
	var a Interger = 1
	if a.Less(2) {
		fmt.Println(a,"less 2")
	}
}
