package main

import "fmt"

func average(xs []float64) (ave float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		ave = 0
	default:
		for _, v := range xs {
			sum += v
		}

		ave = sum / float64(len(xs))
	}
	return
}

func main() {
	a := []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1}
	fmt.Printf("%v", average(a))
}
