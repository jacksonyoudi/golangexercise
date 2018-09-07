package main

import (
	"fmt"
	"time"
)

func main() {
	a := GenAddx(2000)
	CaltimeRun(a(10))
}

func GenAddx(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}

func CaltimeRun(p func(int b) int) {
	start := time.Now()
	p(b)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
