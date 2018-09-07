//package main
//import "fmt"
//func main() {
//	var i1 = 5
//	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
//	var intP *int
//	intP = &i1
//	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
//}

package main

import "fmt"

func main() {
	s := "good bye"
	var p *N = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}
