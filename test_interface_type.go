package main

import "fmt"

type Stringer interface {
	String() string
}

type One struct {
	one Interger
}

func (i *One) String() {
	fmt.Println(i)
}

func main() {
	i := &One{1}
	var v Stringer
	v = i
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}
