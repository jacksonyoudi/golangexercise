package main

import (
	"fmt"
	"math"
	"strconv"
)

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return 0, false
	} // error case
	return math.Sqrt(f), true
}

func main() {
	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
