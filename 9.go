package main

import "fmt"

func main() {
	i := 1
L:
	if i > 10 {
		return
	}
	fmt.Println(i)
	i = i + 1
	goto L
}
