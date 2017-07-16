package main

import (
	"fmt"
)

func main() {
	var a float64
	a = 3.4
	b := int(a)

	// bool类型和int之间无法转换

	fmt.Println(b)

	//定义常量
	const c int = 111

	//定义多个常量
	const a, b, c = 1, "2", 3.23
	const (
		a = "123"
		b = 234
	)
}
