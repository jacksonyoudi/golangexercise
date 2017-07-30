package main

import (
	"fmt"
	"sort"
)

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	sl4 := []int{1,2}
	sl5 := sl4
	fmt.Println(sl5)


	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println(s2)


	m := sort.Ints(sl_from)
	m := sort.SearchInts(a, m)
	fmt.Println(m)
}