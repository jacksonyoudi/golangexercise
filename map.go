package main

import "fmt"

func main() {
	m := map[string]int{'1': 2}

	a, e := m['1']
	fmt.Printf("%s", a)
	fmt.Printf("%s", e)

	b := m['2']
	fmt.Printf(b)

}
